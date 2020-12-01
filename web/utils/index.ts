import { HTMLAttributes } from "react";

export const findByTestAttr = (component: any, attr: any) => component.find(`[data-test='${attr}']`)