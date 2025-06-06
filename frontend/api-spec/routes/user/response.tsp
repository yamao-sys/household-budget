import "../../models/user.tsp";

@summary("User Sign In Response")
model UserSignUpResponse {
  code: integer;
  errors: UserSignUpValidationError;
}

@summary("User Sign In Ok Response")
model UserSignInOkResponse {}

@summary("User Sign In Bad Request Response")
model UserSignInBadRequestResponse {
  errors: string[];
}
