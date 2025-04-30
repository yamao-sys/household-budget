const FOOD = 1;
const DailyGoods = 2;
const Extra = 3;
const Entertainment = 4;
const SelfInvestment = 5;
const Living = 6;
const Miscellaneous = 7;
const TaxSaving = 8;
const Savings = 9;

export const EXPENSE_CATEGORY: { [key: number]: string } = {
  [FOOD]: "食費",
  [DailyGoods]: "日用品",
  [Extra]: "臨時費",
  [Entertainment]: "娯楽費",
  [SelfInvestment]: "自己投資",
  [Living]: "生活費",
  [Miscellaneous]: "諸経費",
  [TaxSaving]: "節税",
  [Savings]: "貯蓄",
};
