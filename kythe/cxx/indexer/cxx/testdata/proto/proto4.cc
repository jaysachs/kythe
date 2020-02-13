#include "kythe/testdata/indexers/proto/testdata4a.pb.h"
#include "kythe/testdata/indexers/proto/testdata4b.pb.h"

//-    P4A=vname("",_, "", "kythe/testdata/indexers/proto/testdata4a.proto","") generates H4A=vname("", _, "bazel-out/bin", "kythe/testdata/indexers/proto/testdata4a.pb.h", "")
//-    P4B=vname("",_, "", "kythe/testdata/indexers/proto/testdata4b.proto","") generates H4B=vname("", _, "bazel-out/bin", "kythe/testdata/indexers/proto/testdata4b.pb.h", "")
//-    P4C=vname("",_, "", "kythe/testdata/indexers/proto/testdata4c.proto","") generates H4C=vname("", _, "bazel-out/bin", "kythe/testdata/indexers/proto/testdata4c.pb.h", "")

// Verify that each .pb.h is generated by exactly the corresponding proto
//- !{ P4A generates H4B }
//- !{ P4A generates H4C }
//- !{ P4B generates H4A }
//- !{ P4B generates H4C }
//- !{ P4C generates H4A }
//- !{ P4C generates H4B }
void fn() {
  using namespace ::pkg::proto4a;

  //- @Thing1 ref Thing1Message?
  //- Thing1Message.node/kind record
  Thing1 thing1;

  using namespace ::pkg::proto4b;

  //- @Thing2 ref Thing2Message?
  //- Thing2Message.node/kind record
  Thing2 thing2;
}
