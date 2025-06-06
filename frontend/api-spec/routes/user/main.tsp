import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request-body.tsp";
import "./response.tsp";

import "../common";

using Http;
using OpenAPI;

@tag("users")
@route("/users")
namespace HouseholdBudgetService.Auth {
  interface ValidateSignUp {
    @route("/validate-sign-up")
    @operationId("post-users-validate-sign-up")
    @summary("User Validate Sign Up")
    @post
    post(
      @body body: UserSignUpInput,
    ): SuccessResponse<UserSignUpResponse>
    | BadRequestErrorResponse<UserSignUpResponse>
    | InternalServerErrorResponse;
  }

  interface SignUp {
    @route("/sign-up")
    @operationId("post-users-sign-up")
    @summary("User Sign Up")
    @post
    post(
      @body body: UserSignUpInput,
    ): SuccessResponse<UserSignUpResponse>
    | BadRequestErrorResponse<UserSignUpResponse>
    | InternalServerErrorResponse;
  }

  interface SignIn {
    @route("/sign-in")
    @operationId("post-users-sign-in")
    @summary("User Sign In")
    @post
    post(
      @body body: UserSignInInput,
    ): SuccessResponseWithSetCookie<UserSignInOkResponse>
    | BadRequestErrorResponse<UserSignInBadRequestResponse>
    | InternalServerErrorResponse;
  }

  interface CheckSignedIn {
    @useAuth([SecuritySchema])
    @route("/check-signed-in")
    @operationId("post-users-check-signed-in")
    @summary("User Check Signed In")
    @get
    get(): SuccessResponse<{ isSignedIn: boolean }>;
  }
}
