<template>
  <v-app-bar app dark elevate-on-scroll scroll-target="#scrolling-techniques-7">
    <v-app-bar-nav-icon></v-app-bar-nav-icon>
    <v-toolbar-title>Monkey's v2</v-toolbar-title>
    <v-spacer></v-spacer>

    <v-btn icon>
      <v-icon>mdi-heart</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>mdi-dots-vertical</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class AppBar extends Vue {}
</script>
