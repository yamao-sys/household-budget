import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;

@service
@doc("household budget")
@OpenAPI.info(#{
  version: "1.0",
  title: "Household Budget Service"
})
@server("https://localhost:8080")
namespace HouseholdBudgetService {};
