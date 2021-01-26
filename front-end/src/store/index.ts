import Vue from "vue";
import Vuex from "vuex";
import { LayoutState } from "./LayoutStore";

Vue.use(Vuex);

export interface RootState {
  layout: LayoutState;
}

export default new Vuex.Store<RootState>({});
