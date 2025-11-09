---
title: "ViajaYa"
category: "App Móvil"
description: "ViajaYa es una aplicación creada para acompañarte en cada aventura. Con ayuda de inteligencia artificial y mapas interactivos, te recomienda lugares, rutas y experiencias únicas según tus intereses. Descubre, planifica y vive cada viaje como nunca antes."
images:
  - "/projects/viajaya.webp"
technologies:
  - "React"
  - "Vite"
  - "Typescript"
  - "Tailwind"
  - "NestJS"
  - "Docker"
repositoryUrl: "https://github.com/orgs/Kodo-Takai/repositories"
liveUrl: "https://frontend-app-kodotakai-eabn.vercel.app/"
order: 1
---

## Características principales

- **Asistente virtual Kodi con IA**: Chatbot con procesamiento de lenguaje natural (NLP) usando Ollama (Qwen2.5:3b-instruct), reconocimiento de voz (STT) con Whisper, síntesis de voz (TTS) y recomendaciones personalizadas basadas en preferencias del usuario

- **Sistema de recomendaciones inteligente**: Genera 3 recomendaciones de destinos basadas en preferencias, presupuesto y categorías favoritas, con integración directa a la base de datos de destinos

- **Gestión de destinos y lugares**: Catálogo de destinos con categorías (hoteles, playas, restaurantes, parques, discotecas, lugares de estudio), búsqueda avanzada y filtrado semántico

- **Agenda de viajes personalizada**: Sistema de agendamiento de destinos con fechas programadas, estados de viaje y gestión de itinerarios

- **Sistema de autenticación completo**: Registro con flujo multi-paso, login con JWT, refresh tokens, recuperación de contraseña y gestión de roles y permisos

- **Perfiles de usuario avanzados**: Perfiles con preferencias de viaje, historial de recomendaciones, interacciones con el chatbot y personalización de experiencias

- **Filtrado inteligente semántico**: Sistema de badges interactivos que analiza descripciones, amenities y reviews para filtrar lugares por relevancia semántica

- **Integración con mapas**: Visualización de destinos en mapas interactivos usando Google Maps API con búsqueda de lugares cercanos

- **PWA (Progressive Web App)**: Aplicación web progresiva con service workers, caché offline, instalación en dispositivos y actualizaciones automáticas

- **Arquitectura modular**: Backend en NestJS con arquitectura hexagonal, frontend en React + TypeScript con Vite, y microservicio de IA en Python/FastAPI

## Rendimiento excepcional

- **Vite como bundler**: Build rápido con HMR en desarrollo y optimización automática de assets (code splitting, tree shaking, minificación)

- **PWA con Workbox**: Caché inteligente de assets estáticos (JS, CSS, HTML, imágenes) hasta 3MB por archivo, funcionamiento offline y actualizaciones en segundo plano

- **React 19 con optimizaciones**: Uso de hooks personalizados, memoización con `useMemo` y `useCallback`, y gestión de estado eficiente con Redux Toolkit y Context API

- **Backend escalable con NestJS**: API REST optimizada con Prisma ORM, transacciones de base de datos, validación de datos con class-validator y documentación automática con Swagger

- **Microservicio de IA optimizado**: Procesamiento asíncrono de NLP, caché de respuestas del chatbot y modelo local (Ollama) para reducir latencia

- **Base de datos PostgreSQL**: Esquema optimizado con Prisma, índices en campos clave, relaciones eficientes y migraciones versionadas

- **Dockerización completa**: Contenedores Docker para frontend (Nginx), backend (Node.js) y servicios de IA, facilitando despliegue y escalabilidad

- **Build de producción optimizado**: TypeScript compilado, assets minificados, lazy loading de rutas y código dividido por chunks para carga progresiva
