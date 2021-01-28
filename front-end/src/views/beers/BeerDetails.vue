<template>
  <BeerDetailsForm />
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import BeerDetailsForm from "@/components/beers/BeerDetailsForm.vue";

@Component({
  components: {
    BeerDetailsForm,
  },
})
export default class BeerDetails extends Vue {}
</script>
