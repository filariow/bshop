<template>
  <v-container>
    <v-card>
      <v-card-title>
        Beers
        <v-spacer></v-spacer>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>

      <v-data-table
        :headers="headers"
        :items="beers"
        :items-per-page="10"
        :search="search"
        :loading="isLoading"
        @click:row="rowClicked"
        class="elevation-1"
        readonly
      ></v-data-table>

      <v-alert prominent type="error" v-if="inError">
        <v-row align="center">
          <v-col class="grow">Error loading beers</v-col>
          <v-col class="shrink">
            <v-btn @click="load" :loading="isLoading">Reload</v-btn>
          </v-col>
        </v-row>
        <v-row
          class="ms-2"
          align="center"
          v-for="err in errors"
          v-bind:key="err"
        >
          <v-col class="grow">{{ err }}</v-col>
        </v-row>
      </v-alert>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { BeersApiService } from "@/services/BeersApiService";
import { Beer } from "@/models/Beers";

@Component
export default class BeersTableForm extends Vue {
  private search = "";

  get headers(): any[] {
    return [
      {
        text: "Name",
        align: "start",
        sortable: true,
        value: "name",
      },
      {
        text: "Price",
        align: "end",
        sortable: true,
        value: "price",
      },
      {
        text: "Size",
        align: "end",
        sortable: true,
        value: "size",
      },
    ];
  }

  private beers: Beer[] = [];
  private errors: any[] = [];
  private isLoading = false;

  get inError(): boolean {
    return this.errors.length > 0;
  }

  async load() {
    try {
      this.isLoading = true;
      this.errors = [];
      this.beers = (await BeersApiService.get<Beer[]>("")).data;
    } catch (e) {
      this.errors.push(e);
    } finally {
      this.isLoading = false;
    }
  }

  async mounted() {
    await this.load();
  }

  rowClicked(item: any): void {
    const id = item.id;
    this.$router.push({ name: "beer-details", params: { id } });
  }
}
</script>
