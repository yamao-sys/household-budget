export const getRequestHeaders = (csrfToken: string) => {
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
    },
  };
};
