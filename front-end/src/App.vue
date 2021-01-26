<template>
  <v-app>
    <app-drawer />
    <app-bar />
    <v-main class="grey lighten-4 fill-height">
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import AppBar from "./components/AppBar.vue";
import AppDrawer from "./components/AppDrawer.vue";

@Component({
  components: {
    AppBar,
    AppDrawer,
  },
})
export default class App extends Vue {}
</script>
