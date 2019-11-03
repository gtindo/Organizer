package main

type FileType struct {
	en, fr, es string
}

// Return fileType value given targeted language
func (ft *FileType) val(lang string) string {
	switch lang {
	case "en":
		return ft.en
	case "fr":
		return ft.fr
	case "es":
		return ft.es
	default:
		return ft.en
	}
}

var AUDIO FileType = FileType{"Audios", "Audios", "Audios"}

var ARCHIVES FileType = FileType{"Archives", "Archives", "Archivo"}

var OTHERS FileType = FileType{"Others", "Autres", "Otras"}

var MEDIA_ARCHIVES = FileType{"Media_archives", "Médias_archivés", "Archivos_de_medios"}

var CAD = FileType{"CAD", "CAD", "CAD"}

var EDA = FileType{"EDA", "EDA", "EDA"}

var DATABASES = FileType{"Databases", "Bases_de_données", "Bases_de_datos"}

var PUBLISHING = FileType{"Publishing", "Édition", "Publicación"}

var DOCUMENTS = FileType{"Documents", "Documents", "Documentos"}

var FINANCIAL_DATAS = FileType{"Financial_datas", "Données_financières", "Datos_financieros"}

var FONTS = FileType{"Fonts", "Polices", "Fuentes"}

var COLORS = FileType{"Colors", "Couleurs", "Colores"}

var PICTURES = FileType{"Pictures", "Images", "Imágenes"}

var VECTOR_GRAPHICS = FileType{"Vector_graphics", "Images_vectorielles", "Gráficos_vectoriales"}

var _3D_GRAPHICS = FileType{"3D_graphics", "Images_3d", "Gráficos_en_3D"}

var SHORTCUTS = FileType{"Shortcuts", "Raccourcis", "Atajos"}

var EXECUTABLES = FileType{"Executables", "Exécutables", "ejecutables"}

var PRESENTATIONS = FileType{"Slides", "Diapositives", "Diapositivas"}

var PLAYLIST = FileType{"Playlists", "Playlists", "Lista_de_reproducción"}

var SOURCES_CODE = FileType{"Sources_code", "Code_sources", "Código fuente"}

var SPREADSHEETS = FileType{"Spreadsheets", "Feuilles_de_calcul", "Hojas_de_cálculo"}

var VIDEOS = FileType{"Videos", "Vidéos", "Videos"}

var VIRTUAL_MACHINES = FileType{"Virtual_machines", "Machines_virtuelles", "Maquinas_virtuales"}

var WEB_PAGES = FileType{"Web_pages", "Pages_web", "Páginas_web"}
