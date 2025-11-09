---
title: "Geekcorp"
category: "Web"
description: "Sitio enfocado en servicios de Call Center y BPO. Ofrece atención multicanal, gestión de cobranzas, agendamiento de citas, soporte técnico remoto, procesamiento de datos, y ventas y encuestas telefónicas. Presenta el equipo, beneficios y casos de clientes. Incluye acceso para clientes (soporte, pagos, intranet) y formulario de contacto. Disponible en español e inglés."
images:
  - "/projects/geekcorp.webp"
technologies:
  - "NextJS"
  - "React"
repositoryUrl: "No disponible"
order: 3
---

## Características principales

- **Internacionalización completa**: Sistema i18n personalizado con soporte para español e inglés, con persistencia en localStorage
- **Diseño responsivo**: Optimizado para móviles, tablets y desktop con Tailwind CSS 4
- **Arquitectura moderna**: Next.js 15 con App Router, React 19 y TypeScript 5
- **Optimización de imágenes**: Soporte para WebP y AVIF, con tamaños adaptativos según dispositivo
- **Componentes modulares**: Estructura organizada con secciones reutilizables (Hero, Services, About, Clients, Contact, Integrations)
- **Tipado fuerte**: TypeScript en todo el proyecto para mayor seguridad y mantenibilidad
- **Navegación inteligente**: Header con menú móvil, dropdown de acceso de clientes y selector de idioma
- **Fuentes personalizadas**: Familia TRIALSagace con múltiples pesos y estilos, optimizada con `font-display: swap`

## Rendimiento excepcional

- **Turbopack en desarrollo**: Compilación rápida con el bundler de próxima generación
- **Renderizado estático**: Exportación estática (`output: export`) para sitios sin servidor, con tiempos de carga mínimos
- **Optimización de imágenes**: Next.js Image con formatos modernos (WebP/AVIF) y tamaños responsivos automáticos
- **Headers de seguridad**: Configuración de seguridad (X-Frame-Options, X-Content-Type-Options, Referrer-Policy)
- **Standalone output**: Configuración para despliegues en Docker y entornos serverless
- **CSS optimizado**: Tailwind CSS 4 con purga automática de estilos no utilizados
- **Carga diferida**: Componentes client-side con `"use client"` solo donde se necesita interactividad
