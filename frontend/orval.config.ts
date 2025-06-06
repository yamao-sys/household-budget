const config = {
  schemas: {
    input: {
      target: "./api-spec/tsp-output/schema/openapi.yaml",
    },
    output: {
      mode: "tags-split",
      client: "react-query",
      httpClient: "fetch",
      schemas: "./apis/model",
      clean: true,
      target: "./apis",
    },
    hooks: {
      afterAllFilesWrite: "prettier --write './apis/**/*.{ts,tsx}'",
    },
  },
};

export default config;
