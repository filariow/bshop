<template>
  <v-container>
    <v-card>
      <v-card-title> {{ beer.name }} </v-card-title>

      <v-row>
        <v-col>
          {{ beer.brewer }}
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          {{ beer.price.toFixed(2) }} €
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          {{ beer.cost.toFixed(2) }} €
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          {{ beer.size }} mL
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { Beer } from "@/models/Beers";
import { Guid } from "guid-typescript";

@Component
export default class BeerDetailsForm extends Vue {
  get beer(): Beer {
    return {
      id: Guid.create(),
      name: "beername",
      brewer: "brewer",
      price: 1.0,
      cost: 1.0,
      size: 33,
    };
  }

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
