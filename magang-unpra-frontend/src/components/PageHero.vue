<template>
  <section class="page-hero">
    <img src="/images/gedung.jpeg" alt="" class="page-hero-bg" />
    <div class="page-hero-overlay"></div>
    <div class="page-hero-content">
      <nav v-if="breadcrumbs && breadcrumbs.length" class="page-hero-breadcrumb">
        <template v-for="(crumb, i) in breadcrumbs" :key="i">
          <RouterLink v-if="crumb.to" :to="crumb.to" class="page-hero-crumb-link">{{ crumb.label }}</RouterLink>
          <span v-else class="page-hero-crumb-current">{{ crumb.label }}</span>
          <span v-if="i < breadcrumbs.length - 1" class="page-hero-crumb-sep">/</span>
        </template>
      </nav>
      <h1 class="page-hero-title">{{ title }}</h1>
      <div class="page-hero-line"></div>
      <p v-if="subtitle" class="page-hero-subtitle">{{ subtitle }}</p>
    </div>
  </section>
</template>

<script setup>
import { RouterLink } from 'vue-router'

defineProps({
  title: { type: String, required: true },
  subtitle: { type: String, default: '' },
  breadcrumbs: { type: Array, default: () => [] },
})
</script>

<style scoped>
.page-hero {
  position: relative;
  width: 100%;
  height: 60vh;
  min-height: 360px;
  display: flex;
  align-items: flex-end;
  overflow: hidden;
}

.page-hero-bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center 40%;
}

.page-hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(105deg, rgba(5, 28, 14, 0.88) 0%, rgba(5, 28, 14, 0.55) 45%, rgba(5, 28, 14, 0.15) 100%);
}

.page-hero-content {
  position: relative;
  z-index: 2;
  padding: 0 72px 64px;
  max-width: 700px;
}

.page-hero-breadcrumb {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
  font-size: 13px;
  opacity: 0;
  animation: heroFadeUp 0.6s 0.1s forwards;
}

.page-hero-crumb-link {
  color: rgba(255, 255, 255, 0.6);
  text-decoration: none;
  transition: color 0.2s;
}

.page-hero-crumb-link:hover {
  color: #fff;
}

.page-hero-crumb-sep {
  color: rgba(255, 255, 255, 0.35);
}

.page-hero-crumb-current {
  color: #fff;
}

.page-hero-title {
  font-size: 2.8rem;
  font-weight: 800;
  color: #fff;
  line-height: 1.15;
  margin: 0 0 12px;
  opacity: 0;
  animation: heroFadeUp 0.6s 0.2s forwards;
}

.page-hero-line {
  width: 0;
  height: 3px;
  background: #5ecb7a;
  border-radius: 2px;
  margin-bottom: 16px;
  animation: heroLineGrow 0.8s 0.4s forwards;
}

.page-hero-subtitle {
  font-size: 0.95rem;
  color: rgba(255, 255, 255, 0.8);
  line-height: 1.7;
  max-width: 520px;
  opacity: 0;
  animation: heroFadeUp 0.6s 0.5s forwards;
}

@keyframes heroFadeUp {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}

@keyframes heroLineGrow {
  from { width: 0; }
  to { width: 56px; }
}

@media (max-width: 960px) {
  .page-hero { height: 50vh; min-height: 300px; }
  .page-hero-content { padding: 0 40px 48px; }
  .page-hero-title { font-size: 2rem; }
}

@media (max-width: 600px) {
  .page-hero { height: 45vh; min-height: 260px; }
  .page-hero-content { padding: 0 20px 36px; }
  .page-hero-title { font-size: 1.5rem; }
}
</style>
