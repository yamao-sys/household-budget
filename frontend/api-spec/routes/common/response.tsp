import "@typespec/http";
import "@typespec/openapi";

using Http;

model SuccessResponse<T> {
  @statusCode status: 200;
  @body body: T
}

model SuccessResponseWithSetCookie<T> {
  @statusCode status: 200;
  @body body: T;
  @header("Set-Cookie") token: string;
}

model BadRequestErrorResponse<T> {
  @statusCode status: 400;
  @body body: T
}

model InternalServerErrorResponse {
  @statusCode status: 500;
  @body body: {
    code: integer;
    message: string;
  }
}
