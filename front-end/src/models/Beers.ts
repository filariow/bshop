import { Guid } from "guid-typescript";

export interface Beer {
    id: Guid,
    name: string;
    brewer: string;
    price: number;
    cost: number;
    size: number;
}