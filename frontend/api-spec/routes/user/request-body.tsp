import "@typespec/http";

import "../common/request.tsp";

using Http;

@summary("User Sign Up Input")
model UserSignUpInput {
  name: string;
  email: string;
  password: string;
}

@summary("User Sign In Input")
model UserSignInInput {
  email: string;
  password: string;
}
