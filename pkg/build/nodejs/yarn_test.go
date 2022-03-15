package nodejs

import (
	"github.com/nais/salsa/pkg/build"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYarnLockParsing(t *testing.T) {

	got, err := YarnDeps(yarnlLockContents)
	assert.NoError(t, err)
	want := map[string]build.Dependency{}
	want["js-tokens"] = build.TestDependency("js-tokens", "4.0.0", "sha512", "45d2547e5704ddc5332a232a420b02bb4e853eef5474824ed1b7986cf84737893a6a9809b627dca02b53f5b7313a9601b690f690233a49bce0e026aeb16fcf29")
	want["loose-envify"] = build.TestDependency("loose-envify", "1.4.0", "sha512", "972bb13c6aff59f86b95e9b608bfd472751cd7372a280226043cee918ed8e45ff242235d928ebe7d12debe5c351e03324b0edfeb5d54218e34f04b71452a0add")
	want["object-assign"] = build.TestDependency("object-assign", "4.1.1", "sha1", "2109adc7965887cfc05cbbd442cac8bfbb360863")
	want["react"] = build.TestDependency("react", "17.0.2", "sha512", "82784fb7be62fddabfcf7ffaabfd1ab0fefc0f4bb9f760f92f5a5deccf0ff9d724e85bbf8c978bea25552b6ddfa6d494663f158dffbeef05c0f1435c94641c6c")
	want["@babel/helper-annotate-as-pure"] = build.TestDependency("@babel/helper-annotate-as-pure", "7.16.7", "sha512", "b3ab76c3f20f4154c0113d478ada28c11197a285fc98282db8fe75f79c03fd024f57ac188e9ba308617b26e30e0d55f664024e7f5193e834fd307119bcb3854b")
	want["@babel/helper-builder-binary-assignment-operator-visitor"] = build.TestDependency("@babel/helper-builder-binary-assignment-operator-visitor", "7.16.7", "sha512", "0ba15d6d16b1623c15bbf81e296e19790d10df501fb6045c7529d9e7f8ec1fa0730896edbd7be1a5f91b9138584aebad640ee7097a4f47a003f73775769acc90")
	want["range-parser"] = build.TestDependency("range-parser", "1.2.1", "sha512", "1eb82cc7ea2baa8ca09e68456ca68713a736f7a27e1d30105e8c4417a80dba944e9a6189468cb37c6ddc700bdea8206bc2bff6cb143905577f1939796a03b04a")
	want["webpack-sources"] = build.TestDependency("webpack-sources", "1.4.3", "sha512", "9604d2dd786fd6508e2a8ed20394e3297323a52338b018cd579faad9ba9eb1b48fb391631a653a8e3b414a45fd6f8a96f3bbc322c0889543ce1216d9acc76379")
	want["util-deprecate"] = build.TestDependency("util-deprecate", "1.0.2", "sha1", "450d4dc9fa70de732762fbd2d4a28981419a0ccf")
	want["unist-util-visit"] = build.TestDependency("unist-util-visit", "2.0.3", "sha512", "889e3f45ccdb24c903d3bd7698692db86a66fd4e01cb815f0e89cbecdffdb404c6205e229a29233ae6a0c8c639ded9d9ab734fe8172696b1e110a01f1968e1ed")
	want["source-map"] = build.TestDependency("source-map", "0.6.1", "sha512", "52381aa6e99695b3219018334fb624739617513e3a17488abbc4865ead1b7303f9773fe1d0f963e9e9c9aa3cf565bab697959aa989eb55bc16396332177178ee")

	build.AssertEqual(t, got, want)
}

const yarnlLockContents = `
# THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
# yarn lockfile v1

"js-tokens@^3.0.0 || ^4.0.0":
  version "4.0.0"
  resolved "https://registry.yarnpkg.com/js-tokens/-/js-tokens-4.0.0.tgz#19203fb59991df98e3a287050d4647cdeaf32499"
  integrity sha512-RdJUflcE3cUzKiMqQgsCu06FPu9UdIJO0beYbPhHN4k6apgJtifcoCtT9bcxOpYBtpD2kCM6Sbzg4CausW/PKQ==

"loose-envify@^1.1.0":
  version "1.4.0"
  resolved "https://registry.yarnpkg.com/loose-envify/-/loose-envify-1.4.0.tgz#71ee51fa7be4caec1a63839f7e682d8132d30caf"
  integrity sha512-lyuxPGr/Wfhrlem2CL/UcnUc1zcqKAImBDzukY7Y5F/yQiNdko6+fRLevlw1HgMySw7f611UIY408EtxRSoK3Q==
  dependencies:
    js-tokens "^3.0.0 || ^4.0.0"

"object-assign@^4.1.1":
  version "4.1.1"
  resolved "https://registry.yarnpkg.com/object-assign/-/object-assign-4.1.1.tgz#2109adc7965887cfc05cbbd442cac8bfbb360863"
  integrity sha1-IQmtx5ZYh8/AXLvUQsrIv7s2CGM=

"@babel/helper-annotate-as-pure@^7.14.5", "@babel/helper-annotate-as-pure@^7.15.4":
  version "7.15.4"
  resolved "https://registry.yarnpkg.com/@babel/helper-annotate-as-pure/-/helper-annotate-as-pure-7.15.4.tgz#3d0e43b00c5e49fdb6c57e421601a7a658d5f835"
  integrity sha512-QwrtdNvUNsPCj2lfNQacsGSQvGX8ee1ttrBrcozUP2Sv/jylewBP/8QFe6ZkBsC8T/GYWonNAWJV4aRR9AL2DA==
  dependencies:
    "@babel/types" "^7.15.4"

"@babel/helper-annotate-as-pure@^7.16.7":
  version "7.16.7"
  resolved "https://registry.yarnpkg.com/@babel/helper-annotate-as-pure/-/helper-annotate-as-pure-7.16.7.tgz#bb2339a7534a9c128e3102024c60760a3a7f3862"
  integrity sha512-s6t2w/IPQVTAET1HitoowRGXooX8mCgtuP5195wD/QJPV6wYjpujCGF7JuMODVX2ZAJOf1GT6DT9MHEZvLOFSw==
  dependencies:
    "@babel/types" "^7.16.7"

"@babel/helper-builder-binary-assignment-operator-visitor@^7.14.5":
  version "7.15.4"
  resolved "https://registry.yarnpkg.com/@babel/helper-builder-binary-assignment-operator-visitor/-/helper-builder-binary-assignment-operator-visitor-7.15.4.tgz#21ad815f609b84ee0e3058676c33cf6d1670525f"
  integrity sha512-P8o7JP2Mzi0SdC6eWr1zF+AEYvrsZa7GSY1lTayjF5XJhVH0kjLYUZPvTMflP7tBgZoe9gIhTa60QwFpqh/E0Q==
  dependencies:
    "@babel/helper-explode-assignable-expression" "^7.15.4"
    "@babel/types" "^7.15.4"

"@babel/helper-builder-binary-assignment-operator-visitor@^7.16.7":
  version "7.16.7"
  resolved "https://registry.yarnpkg.com/@babel/helper-builder-binary-assignment-operator-visitor/-/helper-builder-binary-assignment-operator-visitor-7.16.7.tgz#38d138561ea207f0f69eb1626a418e4f7e6a580b"
  integrity sha512-C6FdbRaxYjwVu/geKW4ZeQ0Q31AftgRcdSnZ5/jsH6BzCJbtvXvhpfkbkThYSuutZA7nCXpPR6AD9zd1dprMkA==
  dependencies:
    "@babel/helper-explode-assignable-expression" "^7.16.7"
    "@babel/types" "^7.16.7"

"@react@>=0.5.1, react@^17.0.2":
  version "17.0.2"
  resolved "https://registry.yarnpkg.com/react/-/react-17.0.2.tgz#d0b5cc516d29eb3eee383f75b62864cfb6800037"
  integrity sha512-gnhPt75i/dq/z3/6q/0asP78D0u592D5L1pd7M8P+dck6Fu/jJeL6iVVK23fptSUZj8Vjf++7wXA8UNclGQcbA==
  dependencies:
    loose-envify "^1.1.0"
    object-assign "^4.1.1"

range-parser@^1.2.1, range-parser@~1.2.1:
  version "1.2.1"
  resolved "https://registry.yarnpkg.com/range-parser/-/range-parser-1.2.1.tgz#3cf37023d199e1c24d1a55b84800c2f3e6468031"
  integrity sha512-Hrgsx+orqoygnmhFbKaHE6c296J+HTAQXoxEF6gNupROmmGJRoyzfG3ccAveqCBrwr/2yxQ5BVd/GTl5agOwSg==

webpack-sources@^1.1.0, webpack-sources@^1.4.3:
  version "1.4.3"
  resolved "https://registry.yarnpkg.com/webpack-sources/-/webpack-sources-1.4.3.tgz#eedd8ec0b928fbf1cbfe994e22d2d890f330a933"
  integrity sha512-lgTS3Xhv1lCOKo7SA5TjKXMjpSM4sBjNV5+q2bqesbSPs5FjGmU6jjtBSkX9b4qW87vDIsCIlUPOEhbZrMdjeQ==
  dependencies:
    source-list-map "^2.0.0"
    source-map "~0.6.1"

util-deprecate@^1.0.1, util-deprecate@^1.0.2, util-deprecate@~1.0.1:
  version "1.0.2"
  resolved "https://registry.yarnpkg.com/util-deprecate/-/util-deprecate-1.0.2.tgz#450d4dc9fa70de732762fbd2d4a28981419a0ccf"
  integrity sha1-RQ1Nyfpw3nMnYvvS1KKJgUGaDM8=

unist-util-visit@2.0.3, unist-util-visit@^2.0.0, unist-util-visit@^2.0.1, unist-util-visit@^2.0.2, unist-util-visit@^2.0.3:
  version "2.0.3"
  resolved "https://registry.yarnpkg.com/unist-util-visit/-/unist-util-visit-2.0.3.tgz#c3703893146df47203bb8a9795af47d7b971208c"
  integrity sha512-iJ4/RczbJMkD0712mGktuGpm/U4By4FfDonL7N/9tATGIF4imikjOuagyMY53tnZq3NP6BcmlrHhEKAfGWjh7Q==
  dependencies:
    "@types/unist" "^2.0.0"
    unist-util-is "^4.0.0"
    unist-util-visit-parents "^3.0.0"

source-map@^0.6.0, source-map@^0.6.1, source-map@~0.6.0, source-map@~0.6.1:
  version "0.6.1"
  resolved "https://registry.yarnpkg.com/source-map/-/source-map-0.6.1.tgz#74722af32e9614e9c287a8d0bbde48b5e2f1a263"
  integrity sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==
`

func TestBuildYarn(t *testing.T) {
	tests := []build.IntegrationTest{
		{
			Name:      "find build file and parse output",
			BuildType: BuildYarn(),
			WorkDir:   "testdata/nodejs/yarn",
			BuildPath: "testdata/nodejs/yarn/yarn.lock",
			Cmd:       "yarn.lock",
			Want: build.Want{
				Key:     "js-tokens",
				Version: "4.0.0",
				Algo:    "sha512",
				Digest:  "45d2547e5704ddc5332a232a420b02bb4e853eef5474824ed1b7986cf84737893a6a9809b627dca02b53f5b7313a9601b690f690233a49bce0e026aeb16fcf29",
			},
		},
		{
			Name:         "cant find build file",
			BuildType:    BuildYarn(),
			WorkDir:      "testdata/whatever",
			Error:        true,
			ErrorMessage: "could not find match, reading dir open testdata/whatever: no such file or directory",
		},
	}

	build.RunTests(t, tests)
}
