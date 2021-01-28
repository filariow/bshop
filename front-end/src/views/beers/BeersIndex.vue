<template>
  <BeersTableForm />
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import BeersTableForm from "@/components/beers/BeersTableForm.vue";

@Component({
  components: {
    BeersTableForm,
  },
})
export default class BeersIndex extends Vue {}
</script>
