import "@typespec/http";

@summary("User Sign Up Validation Error")
model UserSignUpValidationError {
  name?: string[];
  email?: string[];
  password?: string[];
}
