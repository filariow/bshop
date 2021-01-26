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
        class="elevation-1"
        readonly
      ></v-data-table>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class BeersTable extends Vue {
  private search = "";

  get headers() {
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

  private beers = [
    { name: "First Beer", price: "1.0 €", size: "330 mL" },
    { name: "Second Beer", price: "2.0 €", size: "568 mL" },
    { name: "Third Beer", price: "3.0 €", size: "660 mL" },
    { name: "Fourth Beer", price: "4.0 €", size: "750 mL" },
    { name: "Fifth Beer", price: "5.0 €", size: "1000 mL" },
  ];
}
</script>
