<template>
  <nav class="fixed top-0 left-0 right-0 w-full z-50 shadow-sm bg-white">
    <div class="w-full px-10 py-3 flex items-center justify-between">

      <!-- Logo -->
      <RouterLink to="/" class="flex items-center gap-4">
        <img src="/images/logo-telpp.png" alt="TeL Logo" class="h-12 w-auto" />
        <span class="text-[11px] font-bold text-gray-800 tracking-[2px] uppercase leading-snug">
          PT Tanjungenim Lestari<br />Pulp and Paper
        </span>
      </RouterLink>

      <!-- Menu -->
      <div class="flex items-stretch">

        <RouterLink
          to="/"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/') }"
        >
          HOME
        </RouterLink>

        <RouterLink
          to="/our-company"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/our-company') }"
        >
          OUR COMPANY
        </RouterLink>

        <RouterLink
          to="/product"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/product') }"
        >
          PRODUCT
        </RouterLink>

        <RouterLink
          to="/news"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/news') }"
        >
          NEWS
        </RouterLink>

        <!-- Sustainability (dropdown) -->
        <div class="relative group">
          <button
            class="nav-link w-full border-none bg-transparent cursor-pointer"
            :class="{ 'nav-link--active': isActive('/sustainability') }"
          >
            SUSTAINABILITY
            <span style="font-size:9px; margin-top:1px;">▾</span>
          </button>

          <div class="dropdown-menu">
            <RouterLink
              to="/sustainability/forest-management"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/forest-management') }"
            >
              Forest Management
            </RouterLink>
            <RouterLink
              to="/sustainability/people-development"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/people-development') }"
            >
              People Development
            </RouterLink>
            <RouterLink
              to="/sustainability/supply-chain"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/supply-chain') }"
            >
              Supply Chain
            </RouterLink>
            <RouterLink
              to="/sustainability/pulp-process"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/pulp-process') }"
            >
              Pulp Process
            </RouterLink>
            <RouterLink
              to="/sustainability/safety-health"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/safety-health') }"
            >
              Safety &amp; Health
            </RouterLink>

            <div class="relative group/csr">
              <button
                class="dropdown-item w-full text-left border-none bg-transparent cursor-pointer flex justify-between items-center"
                :class="{ 'dropdown-item--active': isActive('/sustainability/csr') }"
              >
                Corporate Social Responsibility
                <span style="color:#aaa; margin-left:8px;">›</span>
              </button>
              <div class="subdropdown-menu">
                <RouterLink
                  to="/sustainability/csr/vision"
                  class="dropdown-item"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/vision') }"
                >
                  Vision And Mission
                </RouterLink>
                <RouterLink
                  to="/sustainability/csr/community"
                  class="dropdown-item"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/community') }"
                >
                  Local Community Development
                </RouterLink>
                <RouterLink
                  to="/sustainability/csr/report"
                  class="dropdown-item"
                  style="border-bottom: none;"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/report') }"
                >
                  CSR Report
                </RouterLink>
              </div>
            </div>
          </div>
        </div>

        <!-- Biodiversity -->
        <a
          href="https://kehati.telpp.com/"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          BIODIVERSITY
        </a>

        <!-- E-Proc -->
        <a
          href="https://eproc.telpp.com/_gst_home.php"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          E-PROC
        </a>

        <!-- E-Recruitment -->
        <a
          href="https://erecruitment.telpp.com/er_tlpp/"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          E-RECRUITMENT
        </a>

        <!-- Contact -->
        <RouterLink
          to="/contact"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/contact') }"
        >
          CONTACT
        </RouterLink>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

const isActive = (path) => {
  if (path === '/') {return route.path === '/'}
  return route.path.startsWith(path)
}
</script>

<style scoped>
/* ── Nav link ── */
.nav-link {
  color: #333333;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1.5px;
  text-transform: uppercase;
  padding: 14px 16px;
  white-space: nowrap;
  display: flex;
  align-items: center;
  gap: 4px;
  text-decoration: none;
  border-bottom: 3px solid transparent;
  transition: color 0.2s, border-color 0.2s;
}

.nav-link:hover {
  color: #5F9E42;
}

.nav-link--active {
  color: #5F9E42;
  border-bottom: 3px solid #5F9E42;
}

/* ── Dropdown utama ── */
.dropdown-menu {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: #ffffff;
  min-width: 230px;
  z-index: 999;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.13);
  border-top: 3px solid #5F9E42;
}

.group:hover > .dropdown-menu {
  display: block;
}

/* ── Sub-dropdown CSR ── */
.subdropdown-menu {
  display: none;
  position: absolute;
  left: 100%;
  top: 0;
  background: #ffffff;
  min-width: 210px;
  z-index: 999;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.13);
  border-top: 3px solid #5F9E42;
}

.group\/csr:hover > .subdropdown-menu {
  display: block;
}

/* ── Item dropdown ── */
.dropdown-item {
  display: block;
  width: 100%;
  padding: 11px 18px;
  font-size: 11px;
  color: #555555;
  border-bottom: 1px solid #f0f0f0;
  text-decoration: none;
  text-align: left;
  transition: background 0.15s, color 0.15s, padding-left 0.15s;
}

.dropdown-item:hover {
  background: #f0f8ea;
  color: #5F9E42;
  padding-left: 22px;
}

.dropdown-item--active {
  background: #f0f8ea;
  color: #5F9E42;
  font-weight: 700;
}
</style>