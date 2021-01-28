<template>
  <v-card>
    <v-card-title> {{ beer.name }} </v-card-title>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field v-model="name" label="Name" read-only></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="brewer"
              label="Brewer"
              read-only
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="price"
              label="Price"
              read-only
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="cost" label="Cost" read-only></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="size" label="Size" read-only></v-text-field>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </v-card>
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

  get name(): string {
    return this.beer.name;
  }

  get brewer(): string {
    return this.beer.brewer;
  }

  get price(): string {
    return `${this.beer.price} €`;
  }

  get cost(): string {
    return `${this.beer.cost} €`;
  }

  get size(): string {
    return `${this.beer.size} mL`;
  }
}
</script>
