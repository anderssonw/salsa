package nodejs

import (
	"reflect"
	"testing"

	"github.com/nais/salsa/pkg/build"
)

func TestYarnLockParsing(t *testing.T) {

	got := YarnDeps(yarnlLockContents)
	want := map[string]build.Dependency{}
	want["js-tokens"] = yarnDep("js-tokens", "4.0.0", "sha512", "RdJUflcE3cUzKiMqQgsCu06FPu9UdIJO0beYbPhHN4k6apgJtifcoCtT9bcxOpYBtpD2kCM6Sbzg4CausW/PKQ==")
	want["loose-envify"] = yarnDep("loose-envify", "1.4.0", "sha512", "lyuxPGr/Wfhrlem2CL/UcnUc1zcqKAImBDzukY7Y5F/yQiNdko6+fRLevlw1HgMySw7f611UIY408EtxRSoK3Q==")
	want["object-assign"] = yarnDep("object-assign", "4.1.1", "sha1", "IQmtx5ZYh8/AXLvUQsrIv7s2CGM=")
	want["react"] = yarnDep("react", "17.0.2", "sha512", "gnhPt75i/dq/z3/6q/0asP78D0u592D5L1pd7M8P+dck6Fu/jJeL6iVVK23fptSUZj8Vjf++7wXA8UNclGQcbA==")
	want["@babel/helper-annotate-as-pure"] = yarnDep("@babel/helper-annotate-as-pure", "7.16.7", "sha512", "s6t2w/IPQVTAET1HitoowRGXooX8mCgtuP5195wD/QJPV6wYjpujCGF7JuMODVX2ZAJOf1GT6DT9MHEZvLOFSw==")
	want["@babel/helper-builder-binary-assignment-operator-visitor"] = yarnDep("@babel/helper-builder-binary-assignment-operator-visitor", "7.16.7", "sha512", "C6FdbRaxYjwVu/geKW4ZeQ0Q31AftgRcdSnZ5/jsH6BzCJbtvXvhpfkbkThYSuutZA7nCXpPR6AD9zd1dprMkA==")
	want["range-parser"] = yarnDep("range-parser", "1.2.1", "sha512", "Hrgsx+orqoygnmhFbKaHE6c296J+HTAQXoxEF6gNupROmmGJRoyzfG3ccAveqCBrwr/2yxQ5BVd/GTl5agOwSg==")
	want["webpack-sources"] = yarnDep("webpack-sources", "1.4.3", "sha512", "lgTS3Xhv1lCOKo7SA5TjKXMjpSM4sBjNV5+q2bqesbSPs5FjGmU6jjtBSkX9b4qW87vDIsCIlUPOEhbZrMdjeQ==")
	want["util-deprecate"] = yarnDep("util-deprecate", "1.0.2", "sha1", "RQ1Nyfpw3nMnYvvS1KKJgUGaDM8=")
	want["unist-util-visit"] = yarnDep("unist-util-visit", "2.0.3", "sha512", "iJ4/RczbJMkD0712mGktuGpm/U4By4FfDonL7N/9tATGIF4imikjOuagyMY53tnZq3NP6BcmlrHhEKAfGWjh7Q==")
	want["source-map"] = yarnDep("source-map", "0.6.1", "sha512", "UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func yarnDep(coordinates, version, alg, digest string) build.Dependency {
	return build.Dependency{
		Coordinates: coordinates,
		Version:     version,
		CheckSum: build.CheckSum{
			Algorithm: alg,
			Digest:    digest,
		},
	}
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
