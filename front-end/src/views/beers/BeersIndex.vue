<template>
  <BeersTable />
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import BeersTable from "@/components/beers/BeersTable.vue";

@Component({
  components: {
    BeersTable,
  },
})
export default class BeersTableIndex extends Vue {}
</script>
