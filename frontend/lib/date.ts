export const getDateString = (date: Date): string => {
  return date.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-");
};

export const getMonthString = (date: Date): string => {
  return date.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit" }).replaceAll("/", "-");
};

export const getMonthLocaleString = (date: Date): string => {
  return `${date.getFullYear()}年${date.getMonth() + 1}月`;
};
