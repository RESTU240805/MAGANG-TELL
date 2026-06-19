import { createRouter, createWebHistory } from 'vue-router'
import OurCompany from '../views/OurCompany.vue'
import OurTeam from '../views/OurTeam.vue'
import ProductView from '../views/ProductView.vue'
import NewsView from '../views/NewsView.vue'
import NewsDetail from '../views/NewsDetail.vue'
import ForestManagementView from '../views/sustainability/ForestManagementView.vue'
import PeopleDevView from '../views/sustainability/PeopleDevView.vue'
import SupplyChainView from '../views/sustainability/SupplyChainView.vue'
import PulpProcessView from '../views/sustainability/PulpProcessView.vue'
import SafetyHealthView from '../views/sustainability/SafetyHealthView.vue'
import VisionView from '../views/sustainability/csr/VisionView.vue'
import CommunityView from '../views/sustainability/csr/CommunityView.vue'
import ReportView from '../views/sustainability/csr/ReportView.vue'
import LoginView from '../views/admin/LoginView.vue'
import DashboardView from '../views/admin/DashboardView.vue'
import NewsListView from '../views/admin/NewsListView.vue'
import NewsFormView from '../views/admin/NewsFormView.vue'
import ProductListView from '../views/admin/ProductListView.vue'
import ProductFormView from '../views/admin/ProductFormView.vue'
import ProductPageEditorView from '../views/admin/ProductPageEditorView.vue'
import ProductDetail from '../views/ProductDetail.vue'
import HomeView from '../views/HomeView.vue'


const router = createRouter({
  history: createWebHistory(),
  routes: [
    // ── General ──
    { path: '/our-company', component: OurCompany },
    { path: '/our-team',    component: OurTeam },
    { path: '/product',     component: ProductView },
    { path: '/news',        component: NewsView },
    { path: '/news/:id',    component: NewsDetail, name: 'NewsDetail' },
    { path: '/contact',     component: () => import('../views/ContactView.vue'), name: 'contact' },

    {path: '/admin/our-company',
  name: 'AdminOurCompany',
  component: () => import('@/views/admin/OurCompanyView.vue'),
},
    // ── Sustainability ──
    { path: '/sustainability/forest-management',  component: ForestManagementView },
    { path: '/sustainability/people-development', component: PeopleDevView },
    { path: '/sustainability/supply-chain',       component: SupplyChainView },
    { path: '/sustainability/pulp-process',       component: PulpProcessView },
    { path: '/sustainability/safety-health',      component: SafetyHealthView },

    // ── CSR ──
    { path: '/sustainability/csr/vision',         component: VisionView },
    { path: '/sustainability/csr/community',      component: CommunityView },
    { path: '/sustainability/csr/report',         component: ReportView },

    // ── Admin ──
    { path: '/admin/login',                       component: LoginView },
    { path: '/admin/dashboard',                   component: DashboardView },
    { path: '/admin/news',                        component: NewsListView },
    { path: '/admin/news/create',                 component: NewsFormView },
    { path: '/admin/news/edit/:id',               component: NewsFormView },
    { path: '/admin/products',                    component: ProductListView },
    { path: '/admin/products/create',             component: ProductFormView },
    { path: '/admin/products/edit/:id',           component: ProductFormView },
    { path: '/admin/slider',                      component: () => import('../views/admin/SliderListView.vue') },
    { path: '/admin/product-page',                component: ProductPageEditorView },
    { path: '/admin/about',                       component: () => import('../views/admin/AboutEditorView.vue') },
    { path: '/admin/team-members',                component: () => import('../views/admin/TeamMemberListView.vue') },
    { path: '/product/:id', component: ProductDetail },

      { path: '/',            component: HomeView },
  ]
})

export default router