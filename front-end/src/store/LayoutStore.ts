import { VuexModule, Module, Mutation, Action, getModule } from 'vuex-module-decorators'
import store from '@/store/index'

export interface LayoutState {
    drawer: boolean;
}

@Module({ namespaced: true, dynamic: true, store, name: 'LayoutStoreModule' })
class LayoutStore extends VuexModule implements LayoutState {
    public drawer = true;

    @Mutation
    public SET_DRAWER(drawer: boolean): void {
        this.drawer = drawer;
    }
    @Action
    public revertDrawer() {
        this.SET_DRAWER(!this.drawer);
    }
}

export const LayoutModule = getModule(LayoutStore);
