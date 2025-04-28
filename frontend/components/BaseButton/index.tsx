import type { FC } from "react";
import { memo } from "react";
import type { JSX } from "react";

type Props = {
  borderColor: string;
  bgColor: string;
  label: string;
} & JSX.IntrinsicElements["button"];

const BaseButton: FC<Props> = memo(function BaseButton({ borderColor, bgColor, label, ...props }: Props) {
  return (
    <>
      <button className={`py-2 px-8 ${borderColor} ${bgColor} rounded-xl text-white`} {...props}>
        {label}
      </button>
    </>
  );
});

export default BaseButton;
