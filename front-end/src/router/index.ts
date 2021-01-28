import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Home from "@/views/Home.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  },
  {
    path: "/help",
    name: "help",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Help.vue")
  },
  {
    path: "/support",
    name: "support",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Support.vue")
  },

  {
    path: "/beers",
    name: "beers",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/beers/BeersIndex.vue")
  },
  {
    path: '/beers/:id',
    name: "beer-details",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/beers/BeerDetails.vue")
  }
];

const router = new VueRouter({
  routes
});

export default router;
