// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maintner

import (
	"bytes"
	"testing"
	"time"
)

var statusTests = []struct {
	msg  string
	want string
}{
	{`Create change

Uploaded patch set 1.

Patch-set: 1 (draft)
Change-id: I38a08cacc17bcd9587475495111fe98f10d6875c
Subject: test: test
Branch: refs/heads/master
Status: draft
Topic:
Commit: fee468c613a70d89f60fb5d683b0f796aabecaac`, "draft"},
	{`Update patch set 1

Change has been successfully cherry-picked as 117ac82c422a11e4dd5f4c14b50bafc1df840481 by Brad Fitzpatrick

Patch-set: 1
Status: merged
Submission-id: 16401-1446004855021-a20b3823`, "merged"},
	{`Create patch set 8

Uploaded patch set 8: Patch Set 7 was rebased.

Patch-set: 8
Subject: devapp: initial support for App Engine Flex
Commit: 17839a9f284b473986f235ad2757a2b445d05068
Tag: autogenerated:gerrit:newPatchSet
Groups: 17839a9f284b473986f235ad2757a2b445d05068`, ""},
}

func TestGetGerritStatus(t *testing.T) {
	for _, tt := range statusTests {
		gc := &GitCommit{Msg: tt.msg}
		got := getGerritStatus(gc)
		if got != tt.want {
			t.Errorf("getGerritStatus msg:\n%s\ngot %s, want %s", tt.msg, got, tt.want)
		}
	}
}

var normalizeTests = []struct {
	in  string
	out string
}{
	{"foo", "foo"},
	{"http://foo", "foo"},
	{"upspin-review.googlesource.com", "upspin.googlesource.com"},
	{"go-review.googlesource.com", "go.googlesource.com"},
}

func TestNormalizeServer(t *testing.T) {
	for _, tt := range normalizeTests {
		got := normalizeGerritServer(tt.in)
		if got != tt.out {
			t.Errorf("normalizeGerritServer(%q) = %q, want %q", tt.in, got, tt.out)
		}
	}
}

func TestGerritProject(t *testing.T) {
	var c Corpus
	c.EnableLeaderMode(new(dummyMutationLogger), "/fake/dir")
	c.TrackGerrit("go.googlesource.com/build")
	gp := c.Gerrit().Project("go-review.googlesource.com", "build")
	if gp == nil {
		t.Errorf("expected go-review.googlesource.com to return a project, got nil")
	}
	gp = c.Gerrit().Project("go-review.googlesource.com", "unknown")
	if gp != nil {
		t.Errorf("expected go-review.googlesource.com to return nil, got a project")
	}
}

var messageTests = []struct {
	in      string
	want    string
	wantNil bool
}{
	{
		in: `Update patch set 1

Patch Set 1: Code-Review+2

Just to confirm, "go test" will consider an empty test file to be passing?

Patch-set: 1
Reviewer: Quentin Smith <13020@62eb7196-b449-3ce5-99f1-c037f21e1705>
Label: Code-Review=+2
`,
		want: "Patch Set 1: Code-Review+2\n\nJust to confirm, \"go test\" will consider an empty test file to be passing?",
	},
	{
		in: `Create change

Uploaded patch set 1: Run-TryBot+1.

Patch-set: 1
Change-id: I1e0035ffba986c3551479d5742809e43da5e7c73
Subject: runtime: fall back to small mmaps if we fail to grow reservation
Branch: refs/heads/master
Status: new
Topic:
Commit: 8776f8d725c001456037e8888a72885d46cd6744
Tag: autogenerated:gerrit:newPatchSet
Groups: 8776f8d725c001456037e8888a72885d46cd6744
Reviewer: Keith Randall <5200@62eb7196-b449-3ce5-99f1-c037f21e1705>
Reviewer: Rick Hudson <5186@62eb7196-b449-3ce5-99f1-c037f21e1705>
Reviewer: Austin Clements <5167@62eb7196-b449-3ce5-99f1-c037f21e1705>
Label: Run-TryBot=+1
Private: false
Work-in-progress: false
`,
		want: "Uploaded patch set 1: Run-TryBot+1.",
	},
	{
		in: `Uploaded patch set 1.

Patch-set: 1
`,
		wantNil: true,
	},
	{
		in: `Create change

Uploaded patch set 1.

Patch-set: 1
Change-id: I3799148a111f1ab6bfee24c9e03e6ebbf9e9595b
Subject: net: make error messages consistent for invalid ports
Branch: refs/heads/master
Commit: 8a7de7048dc194d5e6f761add433b915beebb2e0
Groups: 8a7de7048dc194d5e6f761add433b915beebb2e0
`,
		wantNil: true,
	},
	{
		in: `Create patch set 7

Uploaded patch set 7.: Commit message was updated

Patch-set: 7
Subject: cmd/vet: -lostcancel: check for discarded result of context.WithCancel
Commit: 5487cc78ea332c7b49d43ef5955211387aca73bb
Groups: 5487cc78ea332c7b49d43ef5955211387aca73bb
`,
		wantNil: true,
	},
}

func TestGetGerritMessage(t *testing.T) {
	var c Corpus
	c.EnableLeaderMode(new(dummyMutationLogger), "/fake/dir")
	c.TrackGerrit("go.googlesource.com/build")
	gp := c.gerrit.projects["go.googlesource.com/build"]
	for i, tt := range messageTests {
		gc := &GitCommit{
			Msg:        tt.in,
			CommitTime: time.Now().UTC(),
		}
		msg := gp.getGerritMessage(gc)
		if msg == nil != tt.wantNil {
			if tt.wantNil {
				t.Errorf("%d. getGerritMessage returned item; want nil", i)
			} else {
				t.Errorf("%d. getGerritMessage = nil; want a message", i)
			}
			continue
		}
		if msg == nil {
			continue
		}
		// just checking these get copied through appropriately
		if msg.Version != 1 {
			t.Errorf("%d. getGerritMessage: want Version 1, got %d", i, msg.Version)
		}
		if msg.Date.IsZero() {
			t.Errorf("%d. getGerritMessage: expected Date to be non-zero, got zero", i)
		}
		if msg.Message != tt.want {
			t.Errorf("%d. getGerritMessage = %q; want %q", i, msg.Message, tt.want)
		}
	}
}

func TestOwnerID(t *testing.T) {
	cl := &GerritCL{}
	meta := newGerritMeta(
		&GitCommit{
			Author: &GitPerson{
				Str: "Rick Sanchez <137@62eb7196-b449-3ce5-99f1-c037f21e1705>",
			},
		},
		cl,
	)
	cl.Meta = meta
	cl.Metas = []*GerritMeta{meta}
	cl.Commit = &GitCommit{}

	testCases := []struct {
		cl      *GerritCL
		OwnerID int
	}{
		{&GerritCL{}, -1},
		{cl, 137},
	}
	for _, tc := range testCases {
		if got := tc.cl.OwnerID(); got != tc.OwnerID {
			t.Errorf("cl.OwnerID() = %d; want %d", got, tc.OwnerID)
		}
	}
}

func TestSubject(t *testing.T) {
	cl := &GerritCL{}
	if w, e := cl.Subject(), ""; w != e {
		t.Errorf("cl.Subject() = %q; want %q", w, e)
	}

	testcases := []struct{ msg, subject string }{
		{"maintner: slurp up all the things", "maintner: slurp up all the things"},
		{"cmd/go: build stuff\n\nand do other stuff, too.", "cmd/go: build stuff"},
	}
	for _, tc := range testcases {
		cl = &GerritCL{Commit: &GitCommit{Msg: tc.msg}}
		if cl.Subject() != tc.subject {
			t.Errorf("cl.Subject() = %q; want %q", cl.Subject(), tc.subject)
		}
	}
}

func TestLineValue(t *testing.T) {
	tests := []struct {
		all, prefix, want, wantRest string
	}{
		{
			all:      "foo:  value ",
			prefix:   "foo:",
			want:     "value",
			wantRest: "",
		},
		{
			all:      "foo:  value\n",
			prefix:   "foo:",
			want:     "value",
			wantRest: "",
		},
		{
			all:      "bar: other\nfoo:  value\n",
			prefix:   "foo:",
			want:     "value",
			wantRest: "",
		},
		{
			all:      "notfoo: other\nfoo:  value\n",
			prefix:   "foo:",
			want:     "value",
			wantRest: "",
		},
		{
			all:      "Foo: bar\nLabel: Vote=+1\nLabel: Vote=+2\n",
			prefix:   "Label: ",
			want:     "Vote=+1",
			wantRest: "Label: Vote=+2\n",
		},
		{
			all:      "Label: Vote=+2\n",
			prefix:   "Label: ",
			want:     "Vote=+2",
			wantRest: "",
		},
	}
	for _, tt := range tests {
		got, gotRest := lineValue(tt.all, tt.prefix)
		if got != tt.want {
			t.Errorf("lineValue(%q, %q) returned value %q; want %q", tt.all, tt.prefix, got, tt.want)
		}
		if gotRest != tt.wantRest {
			t.Errorf("lineValue(%q, %q) returned rest %q; want %q", tt.all, tt.prefix, gotRest, tt.wantRest)
		}
	}
}

func TestParseGerritLabelValue(t *testing.T) {
	tests := []struct {
		in        string
		wantLabel string
		wantValue int8
		wantWhose string
	}{
		{"Run-TryBot=+1", "Run-TryBot", 1, ""},
		{"-Run-TryBot", "-Run-TryBot", 0, ""},
		{"-TryBot-Result Gobot Gobot <5976@62eb7196-b449-3ce5-99f1-c037f21e1705>", "-TryBot-Result", 0, "5976@62eb7196-b449-3ce5-99f1-c037f21e1705"},
		{"Run-TryBot=+1 Brad Fitzpatrick <5065@62eb7196-b449-3ce5-99f1-c037f21e1705>", "Run-TryBot", 1, "5065@62eb7196-b449-3ce5-99f1-c037f21e1705"},
		{"TryBot-Result=-1 Gobot Gobot <5976@62eb7196-b449-3ce5-99f1-c037f21e1705>", "TryBot-Result", -1, "5976@62eb7196-b449-3ce5-99f1-c037f21e1705"},
	}
	for _, tt := range tests {
		label, value, whose := parseGerritLabelValue(tt.in)
		if label != tt.wantLabel || value != tt.wantValue || whose != tt.wantWhose {
			t.Errorf("parseGerritLabelValue(%q) = %q, %v, %q; want %q, %v, %q",
				tt.in,
				label, value, whose,
				tt.wantLabel, tt.wantValue, tt.wantWhose)
		}
	}
}

var hashtagTests = []struct {
	commit     string
	wantAdd    string
	wantRemove string
}{
	{
		commit: `Update patch set 1

Hashtag removed:foo

Patch-set: 1
Hashtags:
Tag: autogenerated:gerrit:setHashtag
`,
		wantRemove: "foo",
	},
	{
		commit: `Update patch set 1

Hashtags removed:    foo, bar

Patch-set: 1
Hashtags:
Tag: autogenerated:gerrit:setHashtag
`,
		wantRemove: "foo, bar",
	},
	{
		commit: `Update patch set 1

Hashtag added:   bar

Patch-set: 1
Hashtags:
Tag: autogenerated:gerrit:setHashtag
`,
		wantAdd: "bar",
	},
	{
		commit: `Update patch set 1

Hashtags added: x,y

Patch-set: 1
Hashtags:
Tag: autogenerated:gerrit:setHashtag
`,
		wantAdd: "x,y",
	},
	// No tag:
	{
		commit: `Update patch set 1

Hashtags added: x,y

Patch-set: 1
Hashtags:
Tag: autogenerated:gerrit:otherTag
`,
	},
}

func TestParseHashtags(t *testing.T) {
	for i, tt := range hashtagTests {
		meta := newGerritMeta(&GitCommit{Msg: tt.commit}, nil)
		added, removed, ok := meta.HashtagEdits()
		if ok != (added != "" || removed != "") {
			t.Errorf("%d. inconsistent return values: %q, %q, %v", i, added, removed, ok)
		}
		if string(added) != tt.wantAdd {
			t.Errorf("%d. added = %q; want %q", i, added, tt.wantAdd)
		}
		if string(removed) != tt.wantRemove {
			t.Errorf("%d. removed = %q; want %q", i, removed, tt.wantRemove)
		}

		// And an allocation test too:
		allocs := testing.AllocsPerRun(100, func() {
			_, _, _ = meta.HashtagEdits()
		})
		if allocs > 0 {
			t.Errorf("%d. allocs = %d; want 0", i, int(allocs))
		}
	}
}

// Current tests are a reduction of real CLs in the following format:
// 1 commit is the first commit containing nothing or just some patch notes
// 2 commit adds a tag to CL
// 3 commit removes a tag from the CL
// 4 commit is a commit that is a simple comment or a comment update (this
// is the bug. the fact that the last commit didn't contain any tags to parse.)
var hashtagTestMultipleCommits = []struct {
	meta *GerritMeta
	want GerritHashtags
}{
	{
		meta: &GerritMeta{
			Commit: &GitCommit{
				Msg: `
Update patch set 1

Hashtag added: wait-author ex-wait-release
Hashtag removed: wait-release

Patch-set: 1
Hashtags: wait-author ex-wait-release
Tag: autogenerated:gerrit:setHashtag
`,
				Parents: []*GitCommit{
					{
						Parents: []*GitCommit{
							{
								Parents: []*GitCommit{},
								Msg: `
Update patch set 2

Hashtag added: wait-release

Patch-set: 2
Hashtags: wait-release
Tag: autogenerated:gerrit:setHashtag
`,
							},
						},
						Msg: `
Patch-set: 1
Reviewer: Gobot Gobot <uuid>
Label: TryBot-Result=-1

Update patch set 1

Patch Set 1: Code-Review+1

No hashtags added here test.

Patch-set: 1
Reviewer: Some person <13708@uuid2>
Label: Code-Review=+1
`,
					},
				},
			},
		},
		want: GerritHashtags("wait-author ex-wait-release"),
	},
	{
		meta: &GerritMeta{
			Commit: &GitCommit{
				Msg: `
Create change

Uploaded patch set 1.

Patch-set: 1 (draft)
Change-id: I38a08cacc17bcd9587475495111fe98f10d6875c
Subject: test: test
Branch: refs/heads/master
Status: draft
Topic:
Commit: fee468c613a70d89f60fb5d683b0f796aabecaac
`,
				Parents: []*GitCommit{},
			},
		},
		// No tags added to the CL.
		want: GerritHashtags(""),
	},
}

func TestParseHashtagsWithMultipleCommits(t *testing.T) {
	for _, tt := range hashtagTestMultipleCommits {
		got := tt.meta.Hashtags()
		if tt.want != got {
			t.Errorf("want tags: %s got: %s\n", tt.want, got)
		}
	}
}

var addedSink, removedSink GerritHashtags

func BenchmarkParseHashtags(b *testing.B) {
	b.ReportAllocs()

	var metas []*GerritMeta
	for _, tt := range hashtagTests {
		metas = append(metas, &GerritMeta{Commit: &GitCommit{Msg: tt.commit}})
	}

	for i := 0; i < b.N; i++ {
		for _, meta := range metas {
			addedSink, removedSink, _ = meta.HashtagEdits()
		}
	}
}

func TestGerritHashtagsContains(t *testing.T) {
	tests := []struct {
		set  string
		t    string
		want bool
	}{
		{"", "", false},
		{"x", "", false},
		{"", "x", false},

		{"foo,bar", "foo", true},
		{"foo, bar", "foo", true},
		{" foo, bar", "foo", true},
		{" foo , bar", "foo", true},
		{" foo , bar ", "foo", true},

		{"foo,bar", "bar", true},
		{"foo, bar", "bar", true},
		{" foo, bar", "bar", true},
		{" foo , bar", "bar", true},
		{" foo , bar ", "bar", true},

		{"foo, bar", "fo", false},
		{"foo, bar", "foo, bar", false},
		{"foo, bar", "ba", false},
	}
	for _, tt := range tests {
		got := GerritHashtags(tt.set).Contains(tt.t)
		if got != tt.want {
			t.Errorf("GerritHashtags(%q).Contains(%q) = %v; want %v", tt.set, tt.t, got, tt.want)
		}
	}
}

func TestGerritHashtagsForeach(t *testing.T) {
	tests := []struct {
		set  string
		want string
	}{
		{"", ""},

		{"foo", "foo."},
		{"foo  ", "foo."},
		{"  foo", "foo."},
		{"foo,bar", "foo.bar."},
		{"  foo , bar ", "foo.bar."},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		GerritHashtags(tt.set).Foreach(func(t string) {
			buf.WriteString(t)
			buf.WriteByte('.')
		})
		got := buf.String()
		if got != tt.want {
			t.Errorf("For set %q, got %q; want %q", tt.set, got, tt.want)
		}
	}
}

func TestGerritHashtagsMatch(t *testing.T) {
	tests := []struct {
		set       string
		want      bool // whether "foo" was found
		wantCalls int
	}{
		{"", false, 0},
		{"foo", true, 1},
		{"foo, foo", true, 1},
		{"bar, foo", true, 2},
	}
	for _, tt := range tests {
		calls := 0
		got := GerritHashtags(tt.set).Match(func(t string) bool {
			calls++
			return t == "foo"
		})
		if got != tt.want {
			t.Errorf("For set %q, Match = %v; want %v", tt.set, got, tt.want)
		}
		if calls != tt.wantCalls {
			t.Errorf("For set %q, number of func calls = %v; want %v", tt.set, calls, tt.wantCalls)
		}
	}
}

func TestGerritHashtagsLen(t *testing.T) {
	tests := []struct {
		set  string
		want int
	}{
		{"", 0},
		{"foo", 1},
		{"foo,bar", 2},
		{"foo, bar", 2},
	}
	for _, tt := range tests {
		got := GerritHashtags(tt.set).Len()
		if got != tt.want {
			t.Errorf("For set %q, Len = %v; want %v", tt.set, got, tt.want)
		}
	}
}
