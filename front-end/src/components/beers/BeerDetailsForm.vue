<template>
  <v-card>
    <v-card-title> {{ name }} </v-card-title>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field v-model="name" label="Name" readonly></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="brewer"
              label="Brewer"
              readonly
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="price"
              label="Price"
              readonly
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="cost" label="Cost" readonly></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="size" label="Size" readonly></v-text-field>
          </v-col>
        </v-row>

        <v-alert prominent type="error" v-if="inError">
          <v-row align="center">
            <v-col class="grow">Error loading beer details [id: {{id}}]</v-col>
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
      </v-container>
    </v-form>
  </v-card>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { Beer } from "@/models/Beers";
import { BeersApiService } from "@/services/BeersApiService";

@Component
export default class BeerDetailsForm extends Vue {
  get id(): string {
    return this.$route.params["id"];
  }

  get name(): string {
    return this.beer?.name ?? "";
  }

  get brewer(): string {
    return this.beer?.brewer ?? "";
  }

  get price(): string {
    return this.beer?.price ? `${this.beer?.price} €` : "";
  }

  get cost(): string {
    return this.beer?.cost ? `${this.beer?.cost} €` : "";
  }

  get size(): string {
    return this.beer?.size ? `${this.beer?.size} mL` : "";
  }

  private beer: Beer | null = null;
  private errors: any[] = [];
  private isLoading = false;

  get inError(): boolean {
    return this.errors.length > 0;
  }
  
  async load() {
    try {
      this.isLoading = true;
      this.errors = [];
      this.beer = null;
      this.beer = (await BeersApiService.get<Beer>(`/${this.id}`)).data;
      console.log(this.beer);
    } catch (e) {
      console.log(e);
      this.errors.push(e);
    } finally {
      this.isLoading = false;
    }
  }

  async mounted() {
    await this.load();
  }
}
</script>
