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
        @click:row="rowClicked"
        class="elevation-1"
        readonly
      ></v-data-table>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class BeersTableForm extends Vue {
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
    { id: 1, name: "First Beer", price: "1.0 €", size: "330 mL" },
    { id: 2, name: "Second Beer", price: "2.0 €", size: "568 mL" },
    { id: 3, name: "Third Beer", price: "3.0 €", size: "660 mL" },
    { id: 4, name: "Fourth Beer", price: "4.0 €", size: "750 mL" },
    { id: 5, name: "Fifth Beer", price: "5.0 €", size: "1000 mL" },
  ];
  
  rowClicked(item: any): void {
    const id = item.id;
		this.$router.push({ name: "beer-details", params: { id } });
	}
}
</script>
