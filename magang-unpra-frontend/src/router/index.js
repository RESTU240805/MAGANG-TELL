import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProductView from '../views/ProductView.vue'
import NewsView from '../views/NewsView.vue'
import ForestManagementView from '../views/sustainability/ForestManagementView.vue'
import PeopleDevView from '../views/sustainability/PeopleDevView.vue'
import SupplyChainView from '../views/sustainability/SupplyChainView.vue'
import PulpProcessView from '../views/sustainability/PulpProcessView.vue'
import SafetyHealthView from '../views/sustainability/SafetyHealthView.vue'
import CSRView from '../views/sustainability/CSRView.vue'
import VisionView from '../views/sustainability/csr/VisionView.vue'
import CommunityView from '../views/sustainability/csr/CommunityView.vue'
import ReportView from '../views/sustainability/csr/ReportView.vue'
import LoginView from '../views/admin/LoginView.vue'
import DashboardView from '../views/admin/DashboardView.vue'


const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: HomeView },
    { path: '/product', component: ProductView },
    { path: '/news', component: NewsView },
    { path: '/sustainability/forest-management', component: ForestManagementView },
    { path: '/sustainability/people-development', component: PeopleDevView },
    { path: '/sustainability/supply-chain', component: SupplyChainView },
    { path: '/sustainability/pulp-process', component: PulpProcessView },
    { path: '/sustainability/safety-health', component: SafetyHealthView },
    { path: '/sustainability/csr', component: CSRView },
    { path: '/sustainability/csr/vision', component: VisionView },
{ path: '/sustainability/csr/community', component: CommunityView },
{ path: '/sustainability/csr/report', component: ReportView },
{ path: '/admin/login', component: LoginView },
{ path: '/admin/dashboard', component: DashboardView },
  ]
})

export default router