<template>
    <v-container>
        <v-row class="text-center">
            <v-col class="mb-4">
                <h1 class="display-2 font-weight-bold mb-3">This pretends to be an Home Page!</h1>
                <!-- <v-col cols="12">
                    <v-img src="../../assets/logo.svg" class="my-3" height="200" contain/>
                </v-col> -->
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts">
    import { Vue, Component } from "vue-property-decorator";
    @Component({})
    export default class Home extends Vue {}
</script>
