<template>
  <v-container>
    <v-navigation-drawer v-model="drawer" app dark expand-on-hover>
      <template v-slot:append>
        <v-list dense nav>
          <v-list-item
            v-for="item in footer"
            :key="item.title"
            exact
            link
            :to="{ name: item.page }"
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </template>

      <v-list-item />
      <v-divider />
      <v-list dense nav>
        <v-list-item
          v-for="item in header"
          :key="item.title"
          exact
          link
          :to="{ name: item.page }"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <v-divider />
      <v-list dense nav>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          exact
          link
          :to="{ name: item.page }"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </v-container>
</template>

<script lang="ts">
import { LayoutModule } from "@/store/LayoutStore";
import { Component, Vue } from "vue-property-decorator";

@Component
export default class AppDrawer extends Vue {
  get drawer(): boolean {
    return LayoutModule.drawer;
  }
  set drawer(value: boolean) {
    LayoutModule.SET_DRAWER(value);
  }

  get header() {
    return [
      {
        title: "Home",
        icon: "mdi-home",
        page: "Home",
      },
    ];
  }
  get items() {
    return [
      {
        title: "Beers",
        icon: "mdi-glass-mug-variant",
        page: "Beers",
      },
    ];
  }

  get footer() {
    return [
      {
        title: "About",
        icon: "mdi-information",
        page: "About",
      },
      {
        title: "Help",
        icon: "mdi-help",
        page: "Help",
      },
      {
        title: "Support",
        icon: "mdi-heart",
        page: "Support",
      },
    ];
  }
}
</script>
