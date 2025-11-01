# loadept-mcp-filesystem

**Versión:** 0.3.3  
**Lenguaje:** Python 3.13+  
**Tipo:** Paquete Python

## Descripción
Servidor MCP para operaciones seguras del sistema de archivos.

## Ejecución
```bash
uvx loadept-mcp-filesystem
```

## Variables de Entorno
```bash
BASE_PATH="/ruta/al/directorio/base"  # Obligatorio para seguridad
```

## Herramientas
- `list_directory`: Lista contenido de directorio
- `find_results`: Busca archivos por patrón
- `read_content`: Lee contenido de archivo
- `write_content`: Escribe archivo
- `open_file`: Abre archivo con aplicación predeterminada
- `create_directory`: Crea directorio

## Instalación
```bash
# Instalar UV si no lo tienes
curl -LsSf https://astral.sh/uv/install.sh | sh  # macOS/Linux
# o
powershell -c "irm https://astral.sh/uv/install.ps1 | iex"  # Windows

# Ejecutar
uvx loadept-mcp-filesystem
```
