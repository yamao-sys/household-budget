import { test, expect } from "@playwright/test";

test.describe("/monthly_budget", () => {
  test("当月操作", async ({ page }) => {
    await page.goto("/sign_in");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_in_1@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログインする" }).click();
    await page.waitForURL("/monthly_budget");

    const now = new Date();
    const monthString = now.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit" }).replaceAll("/", "-");
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥35,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,065,000", { exact: true })).toBeVisible();

    // NOTE: user1の当月の支出詳細画面でカテゴリ毎の支出合計が表示されること
    await page.getByRole("link", { name: "支出詳細へ" }).click();
    await page.waitForURL(`/monthly_budget/${monthString}`);
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible(); // 食費
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible(); // 日用品
    await expect(page.getByText("¥20000", { exact: true })).toBeVisible(); // 娯楽費
    await expect(page.getByText("¥35000", { exact: true })).toBeVisible(); // 支出合計

    await page.getByRole("link", { name: "Household Budget" }).click();

    // NOTE: user1の当月の収入が表示されること
    await expect(page.getByText("収入: ¥1100000", { exact: true })).toBeVisible();
    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社1" })).toBeVisible();
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがある日付で支出合計が表示されること(10日目)
    await expect(page.getByText("支出: ¥15000", { exact: true })).toBeVisible();
    // NOTE: データがある日付でdialog devで支出が表示されること
    await page.getByText("10日", { exact: true }).nth(0).click();
    const modal = page.locator('div[role="dialog"]');
    await expect(modal).toBeVisible();
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "食費" })).toBeVisible();
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "日用品" })).toBeVisible();

    // NOTE: データがある日付でdialog devで支出が登録できること
    await page.getByLabel("支出金額").fill("3000");
    await page.getByLabel("適用").fill("映画");
    await page.getByLabel("カテゴリ").selectOption({ label: "娯楽費" });
    await page.getByRole("button", { name: "支出を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥38,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,062,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();

    await page.getByText("10日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "食費" })).toBeVisible();
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "日用品" })).toBeVisible();
    await expect(page.getByText("¥3000", { exact: true })).toBeVisible();
    await expect(page.getByText("映画", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "娯楽費" })).toBeVisible();

    // NOTE: モーダルを閉じる
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがない日付でdialog devで支出が登録できること(バリデーションエラーがあれば表示されること)
    await page.getByText("15日", { exact: true }).nth(0).click();
    await page.getByRole("button", { name: "支出を登録する" }).click();
    // NOTE: バリデーションメッセージが表示されることを確認
    await expect(page.getByText("金額は必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("カテゴリは必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("適用は必須入力です。", { exact: true })).toBeVisible();

    await page.getByLabel("支出金額").fill("2000");
    await page.getByLabel("適用").fill("書籍");
    await page.getByLabel("カテゴリ").selectOption({ label: "自己投資" });
    await page.getByRole("button", { name: "支出を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,060,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();

    await page.getByText("15日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("書籍", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "自己投資" })).toBeVisible();

    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: 登録した支出が当月の支出詳細画面でカテゴリ毎の支出合計に反映されること
    await page.getByRole("link", { name: "支出詳細へ" }).click();
    await page.waitForURL(`/monthly_budget/${monthString}`);
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible(); // 食費
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible(); // 日用品
    await expect(page.getByText("¥23000", { exact: true })).toBeVisible(); // 娯楽費
    await expect(page.getByText("¥2000", { exact: true })).toBeVisible(); // 自己投資
    await expect(page.getByText("¥40000", { exact: true })).toBeVisible(); // 支出合計

    await page.getByRole("link", { name: "Household Budget" }).click();
    await page.waitForURL(`/monthly_budget`);

    await page.locator(".fc-icon-chevron-left").click(); // 先月に遷移
    await expect(page.getByText("収入合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥0", { exact: true })).toBeVisible();

    /*** 収入登録 ***/
    await page.locator(".fc-icon-chevron-right").click(); // 当月に遷移

    // // NOTE: データがある日付でdialog devで支出が登録できること
    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.locator('div[role="dialog"]')).toBeVisible();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社1" })).toBeVisible();

    await page.getByLabel("収入金額").fill("300000");
    await page.getByLabel("顧客名").fill("テスト収入追加株式会社1");
    await page.getByRole("button", { name: "収入を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,400,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,360,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥1400000", { exact: true })).toBeVisible();

    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社1" })).toBeVisible();
    await expect(page.getByText("¥300000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト収入追加株式会社1" })).toBeVisible();

    // NOTE: モーダルを閉じる
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがない日付でdialog devで収入が登録できること(バリデーションエラーがあれば表示されること)
    await page.getByText("15日", { exact: true }).nth(0).click();
    await page.getByRole("button", { name: "収入を登録する" }).click();
    // NOTE: バリデーションメッセージが表示されることを確認
    await expect(page.getByText("金額は必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("顧客名は必須入力です。", { exact: true })).toBeVisible();

    await page.getByLabel("収入金額").fill("200000");
    await page.getByLabel("顧客名").fill("テスト収入追加株式会社2");
    await page.getByRole("button", { name: "収入を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,600,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,560,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥1400000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥200000", { exact: true })).toBeVisible();

    await page.getByText("15日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥200000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト収入追加株式会社2" })).toBeVisible();

    await page.getByRole("button", { name: "閉じる" }).click();
  });

  test("先月操作", async ({ page }) => {
    await page.goto("/sign_in");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_in_2@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログインする" }).click();
    await page.waitForURL("/monthly_budget");

    const now = new Date();
    now.setMonth(now.getMonth() - 1);
    const monthString = now.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit" }).replaceAll("/", "-");

    // NOTE: 支出合計が表示されること
    // NOTE: user2の当月
    await expect(page.getByText("収入合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥0", { exact: true })).toBeVisible();
    // NOTE: user2の先月
    await page.locator(".fc-icon-chevron-left").click();
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥35,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,065,000", { exact: true })).toBeVisible();

    // NOTE: user2の先月の支出詳細画面でカテゴリ毎の支出合計が表示されること
    await page.getByRole("link", { name: "支出詳細へ" }).click();
    await page.waitForURL(`/monthly_budget/${monthString}`);
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible(); // 食費
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible(); // 日用品
    await expect(page.getByText("¥20000", { exact: true })).toBeVisible(); // 娯楽費
    await expect(page.getByText("¥35000", { exact: true })).toBeVisible(); // 支出合計

    await page.getByRole("link", { name: "Household Budget" }).click();
    await page.locator(".fc-icon-chevron-left").click(); // 先月に戻る

    // NOTE: user2の先月の収入が表示されること
    await expect(page.getByText("収入: ¥1100000", { exact: true })).toBeVisible();
    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社2" })).toBeVisible();
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがある日付で支出合計が表示されること(10日目)
    await expect(page.getByText("支出: ¥15000", { exact: true })).toBeVisible();
    // NOTE: データがある日付でdialog devで支出が表示されること
    await page.getByText("10日", { exact: true }).nth(0).click();
    const modal = page.locator('div[role="dialog"]');
    await expect(modal).toBeVisible();
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "食費" })).toBeVisible();
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "日用品" })).toBeVisible();

    // NOTE: データがある日付でdialog devで支出が登録できること
    await page.getByLabel("支出金額").fill("3000");
    await page.getByLabel("適用").fill("映画");
    await page.getByLabel("カテゴリ").selectOption({ label: "娯楽費" });
    await page.getByRole("button", { name: "支出を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥38,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,062,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();

    await page.getByText("10日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "食費" })).toBeVisible();
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "日用品" })).toBeVisible();
    await expect(page.getByText("¥3000", { exact: true })).toBeVisible();
    await expect(page.getByText("映画", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "娯楽費" })).toBeVisible();

    // NOTE: モーダルを閉じる
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがない日付でdialog devで支出が登録できること(バリデーションエラーがあれば表示されること)
    await page.getByText("15日", { exact: true }).nth(0).click();
    await page.getByRole("button", { name: "支出を登録する" }).click();
    // NOTE: バリデーションメッセージが表示されることを確認
    await expect(page.getByText("金額は必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("カテゴリは必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("適用は必須入力です。", { exact: true })).toBeVisible();

    await page.getByLabel("支出金額").fill("2000");
    await page.getByLabel("適用").fill("書籍");
    await page.getByLabel("カテゴリ").selectOption({ label: "自己投資" });
    await page.getByRole("button", { name: "支出を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,100,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,060,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();

    await page.getByText("15日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("書籍", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "自己投資" })).toBeVisible();

    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: 登録した支出が先月の支出詳細画面でカテゴリ毎の支出合計に反映されること
    await page.getByRole("link", { name: "支出詳細へ" }).click();
    await page.waitForURL(`/monthly_budget/${monthString}`);
    await expect(page.getByText("¥10000", { exact: true })).toBeVisible(); // 食費
    await expect(page.getByText("¥5000", { exact: true })).toBeVisible(); // 日用品
    await expect(page.getByText("¥23000", { exact: true })).toBeVisible(); // 娯楽費
    await expect(page.getByText("¥2000", { exact: true })).toBeVisible(); // 自己投資
    await expect(page.getByText("¥40000", { exact: true })).toBeVisible(); // 支出合計

    await page.getByRole("link", { name: "Household Budget" }).click();
    await page.waitForURL(`/monthly_budget`);
    await page.locator(".fc-icon-chevron-left").click(); // 先月に戻る

    await page.locator(".fc-icon-chevron-right").click(); // 当月に遷移
    await expect(page.getByText("収入合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥0", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥0", { exact: true })).toBeVisible();

    /*** 収入登録 ***/
    await page.locator(".fc-icon-chevron-left").click(); // 先月に遷移

    // NOTE: データがある日付でdialog devで支出が登録できること
    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.locator('div[role="dialog"]')).toBeVisible();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社2" })).toBeVisible();

    await page.getByLabel("収入金額").fill("300000");
    await page.getByLabel("顧客名").fill("テスト収入追加株式会社1");
    await page.getByRole("button", { name: "収入を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,400,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,360,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥20000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥1400000", { exact: true })).toBeVisible();

    await page.getByText("20日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥1100000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト株式会社2" })).toBeVisible();
    await expect(page.getByText("¥300000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト収入追加株式会社1" })).toBeVisible();

    // NOTE: モーダルを閉じる
    await page.getByRole("button", { name: "閉じる" }).click();

    // NOTE: データがない日付でdialog devで収入が登録できること(バリデーションエラーがあれば表示されること)
    await page.getByText("15日", { exact: true }).nth(0).click();
    await page.getByRole("button", { name: "収入を登録する" }).click();
    // NOTE: バリデーションメッセージが表示されることを確認
    await expect(page.getByText("金額は必須入力です。", { exact: true })).toBeVisible();
    await expect(page.getByText("顧客名は必須入力です。", { exact: true })).toBeVisible();

    await page.getByLabel("収入金額").fill("200000");
    await page.getByLabel("顧客名").fill("テスト収入追加株式会社2");
    await page.getByRole("button", { name: "収入を登録する" }).click();

    // NOTE: 結果が支出と利益に反映されること
    await expect(page.getByText("収入合計: ¥1,600,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出合計: ¥40,000", { exact: true })).toBeVisible();
    await expect(page.getByText("利益: ¥1,560,000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥18000", { exact: true })).toBeVisible();
    await expect(page.getByText("支出: ¥2000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥1400000", { exact: true })).toBeVisible();
    await expect(page.getByText("収入: ¥200000", { exact: true })).toBeVisible();

    await page.getByText("15日", { exact: true }).nth(0).click();
    await expect(page.getByText("¥200000", { exact: true })).toBeVisible();
    await expect(page.getByRole("cell", { name: "テスト収入追加株式会社2" })).toBeVisible();

    await page.getByRole("button", { name: "閉じる" }).click();
  });
});
