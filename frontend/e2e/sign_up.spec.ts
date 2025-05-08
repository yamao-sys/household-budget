import { test, expect } from "@playwright/test";

test.describe("/sign_up", () => {
  test("正常系", async ({ page }) => {
    await page.goto("/sign_up");

    // NOTE: 会員登録フォームを入力
    await page.getByRole("textbox", { name: "ユーザ名" }).fill("test_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_up_1@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "登録する" }).click();

    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("会員登録が完了しました");
      await dialog.accept();
    });
    await page.waitForURL("/sign_in");
    await expect(page).toHaveURL("/sign_in");
  });

  test("異常系", async ({ page }) => {
    await page.goto("/sign_up");

    page.on("console", (msg) => console.log(msg.text()));

    // NOTE: 会員登録フォームを入力
    await page.getByRole("button", { name: "登録する" }).click();

    // NOTE: バリデーションエラーが表示されること
    await expect(page.getByText("ユーザ名は必須入力です。")).toBeVisible();
    await expect(page.getByText("Emailは必須入力です。")).toBeVisible();
    await expect(page.getByText("パスワードは必須入力です。")).toBeVisible();

    // 入力し直して登録できること
    await page.getByRole("textbox", { name: "ユーザ名" }).fill("test_name");
    await page.getByRole("textbox", { name: "Email" }).fill("test_sign_up_2@example.com");
    await page.getByRole("textbox", { name: "パスワード" }).fill("password");
    await page.getByRole("button", { name: "登録する" }).click();

    page.on("dialog", async (dialog) => {
      expect(dialog.message()).toContain("会員登録が完了しました");
      await dialog.accept();
    });
    await page.waitForURL("/sign_in");
    await expect(page).toHaveURL("/sign_in");
  });
});
