package DelTree

import (
  // "isabelle/exported/HOL"
  . "isabelle/exported/Tree"
)


func Rightest[a any] (x0 Tree[a]) a {
  {
    q, m := x0.(Node[a]);
    if m {
      _, xa, c := Node_dest(q);
      if c == (Tree[a](Tip[a]{})) {
        return xa;
      }
    }
  };
  {
    q, m := x0.(Node[a]);
    if m {
      _, _, p := Node_dest(q);
      q, m := p.(Node[a]);
      if m {
        vc, vaa, vba := Node_dest(q);
        return Rightest[a](Tree[a](Node[a]{vc, vaa, vba}));
      }
    }
  };
  panic("match failed");
}

