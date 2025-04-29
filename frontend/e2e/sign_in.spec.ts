import { test, expect } from "@playwright/test";

test.describe("/sign_in", () => {
  test("正常系", async ({ page }) => {
    await page.goto("/sign_in");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_in_1@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログインする" }).click();

    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("ログインしました");
      await dialog.accept();
    });
    await page.waitForURL("/");
    await expect(page).toHaveURL("/");
  });

  test("異常系", async ({ page }) => {
    await page.goto("/sign_in");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_in_@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログインする" }).click();

    // NOTE: バリデーションエラーが表示されること
    await expect(page.getByText("メールアドレスまたはパスワードが正しくありません")).toBeVisible();

    // 入力し直して登録できること
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_in_2@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "ログインする" }).click();

    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("ログインしました");
      await dialog.accept();
    });
    await page.waitForURL("/");
    await expect(page).toHaveURL("/");
  });
});
