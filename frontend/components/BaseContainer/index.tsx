import type { ReactNode } from "react";

type Props = {
  children: ReactNode;
  containerWidth?: string;
};

const BaseContainer = ({ children, containerWidth = "w-4/5" }: Props) => {
  return (
    <div className={`${containerWidth} mx-auto`}>
      <div>{children}</div>
    </div>
  );
};

export default BaseContainer;
