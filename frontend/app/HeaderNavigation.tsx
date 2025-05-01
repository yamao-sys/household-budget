import type { FC } from "react";
import { Link } from "react-router";
import { NAVIGATION_PAGE_LIST } from "./routes";

type Props = {
  isSignedIn: boolean;
  children: React.ReactNode;
};

export const HeaderNavigation: FC<Props> = ({ isSignedIn, children }: Props) => {
  return (
    <>
      <header className='bg-white py-4 px-6 border-b fixed top-0 left-0 w-full z-100'>
        <div className='mx-auto flex flex-col md:flex-row md:items-center md:justify-between'>
          <h1 className='text-center text-lg md:text-2xl font-semibold text-gray-800'>
            <Link to={NAVIGATION_PAGE_LIST.monthlyBudgetPage}>Household Budget</Link>
          </h1>

          <nav className='mt-2 md:mt-0 flex justify-center md:justify-end items-center space-x-4 text-sm md:text-base'>
            {isSignedIn ? (
              <Link to={NAVIGATION_PAGE_LIST.monthlyBudgetPage} className='underline text-gray-600 hover:text-gray-900 transition'>
                月の家計簿
              </Link>
            ) : (
              <Link to={NAVIGATION_PAGE_LIST.signInPage} className='underline text-gray-600 hover:text-gray-900 transition'>
                ログイン
              </Link>
            )}
            <Link to={NAVIGATION_PAGE_LIST.signUpPage} className='underline text-gray-600 hover:text-gray-900 transition'>
              会員登録
            </Link>
          </nav>
        </div>
      </header>
      <div className='mx-auto pt-20 px-6'>{children}</div>
    </>
  );
};
