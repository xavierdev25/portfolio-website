---
title: "JorStock"
category: "App de Escritorio"
description: "JorStock es una aplicación de escritorio para gestionar inventario de autopartes. Permite registrar, editar y eliminar productos; buscar por nombre y proveedor con coincidencia parcial; ordenar por nombre, stock, precio y fecha; y visualizar agrupaciones por proveedor con estilos alternados. Gestiona proveedores automáticamente y valida datos (precio/stock)."
images:
  - "/projects/jorstock.webp"
technologies:
  - "C#"
  - ".NET Framework"
  - "Windows Forms"
  - "MongoDB"
repositoryUrl: "https://github.com/xavierdev25/jorstock-app.git"
liveUrl: "No disponible"
order: 4
---

## Características principales

- **CRUD de productos con MongoDB**: creación, edición, búsqueda y eliminación con validaciones
- **Gestión automática de proveedores**: alta implícita y generación de supplierId único
- **Interfaz WinForms moderna**: bordes redondeados y placeholders inteligentes en entradas
- **Búsqueda combinada**: por nombre de autoparte y proveedor mediante expresiones regulares
- **Ordenamiento dinámico**: por nombre, stock, precio y fecha desde menú contextual
- **Agrupación visual por proveedor**: colores alternados por grupo en el DataGridView

## Rendimiento excepcional

- **Operaciones asíncronas con MongoDB**: evita bloqueo de la interfaz usando métodos async/await
- **Reducción de round-trips**: construye filtros y hace lookups de proveedor solo cuando se requieren
- **Ordenamiento en memoria**: aplica DataView.Sort sin reconsultar la base de datos
- **DataGridView de solo lectura**: minimiza costos de edición y mejora la velocidad de renderizado
- **Validaciones tempranas**: detiene operaciones inválidas antes de golpear la base (precio, stock, placeholders)
- **Control de concurrencia**: flag cargandoProductos impide ejecuciones simultáneas de carga
- **Actualizaciones selectivas**: UpdateOne con sets puntuales reduce operaciones de escritura y locks
- **Preparado para índices**: campos consultados frecuentemente (nombre, codigo_proveedor, serial) listos para indexación
- **Agrupación eficiente de proveedores**: usa Dictionary para asignación de grupos en una pasada
- **Carga diferida de nombres de proveedor**: solo resolve nombre cuando se arma cada fila necesaria
