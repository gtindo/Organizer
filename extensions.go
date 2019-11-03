package main

var ARCHIVES_EXT = []string{
	"?q?", "7z", "aapkg", "ace", "alz", "appx", "at3", "bke", "arc", "arj", "ass",
	"b", "ba", "big", "bin", "bsjn", "bkf", "bzip2", "bld", "cab", "c4", "cals", "clipflair", "cpt", "daa",
	"deb", "ddz", "dn", "dpe", "egg", "egt", "ecab", "esd", "ess", "flipchart", "gbp", "gho",
	"gzip", "ipg", "jar", "lbr", "lqr", "lha", "lzip", "lzo", "lzma", "lzx", "mbw", "mhtml", "mpq",
	"nth", "oar", "osz", "pak", "par", "paf", "pea", "pyk", "pk3", "pk4", "rar", "rag", "rags",
	"rax", "rpm", "sb", "sb2", "sb", "sb2", "sen", "sit", "sis", "sisx", "skb", "sq", "swm", "szs", "tar",
	"gz", "tb", "tib", "uha", "uue", "viv", "vol", "vsa", "wax", "wim", "xz", "z", "zoo", "zip",
}

var MEDIA_ARCHIVES_EXT = []string{
	"iso", "nrg", "img", "adf", "adz", "dms", "dsk", "d64", "sdi", "mds", "mdx", "dmg", "cdi", "cue", "cif",
	"c2d", "daa", "b6t",
}

var COMPUTER_AID_DESIGN_EXT = []string{
	"3dxml", "3mf", "acp", "amf", "aec", "ar", "art", "asc", "asm", "bim", "brep", "c3d", "ccc", "ccm", "ccs", "cad",
	"catdrawing", "catpart", "catproduct", "catprocess", "cgr", "ckd", "ckt", "co", "drw", "dft", "dgn", "dgk", "dmt", "dxf", "dwb",
	"dwf", "dwg", "easm", "edrw", "emb", "eprt", "escpcb", "escsch", "esw", "excellon", "exp", "f3d", "fcstd", "fm", "fmz", "g", "gbr",
	"glm", "grb", "gtc", "iam", "icd", "idw", "ifc", "iges", "intergraph", "ipn", "ipt", "jt", "mcd", "model", "ocd", "par", "pipe",
	"pln", "prt", "psm", "psmodel", "pwi", "pyt", "skp", "rlf", "rvm", "rvt", "rfa", "s12", "scad", "scdoc", "sldasm", "slddrw", "sldprt",
	"dotxsi", "step", "stl", "std", "tct", "tcw", "unv", "vc6", "vlm", "vs", "wrl", "x_b", "x_t", "xe", "zofzproj",
}

var ELECTRONIC_DESIGN_AUTOMATION_EXT = []string{
	"brd", "bsdl", "cdl", "cpf", "def", "dspf", "edif", "fsdb", "gdsii", "hex", "lef", "lib", "ms12", "oasis", "openacces",
	"psf", "psfxl", "sdc", "sdf", "spef", "spi", "ci", "pt", "srec", "sst2", "stil", "sv", "s*p", "upf", "v", "vcd", "vhd", "vhdl", "wgl",
}

var DATABASE_EXT = []string{
	"4db", "4dd", "4dindy", "4dindx", "4dr", "accdb", "accde", "adt", "apr", "box", "chml", "daf", "dat", "db", "dbf", "dta", "egt",
	"ess", "eap", "fdb", "fdb", "fp", "fp3", "fp5", "fp7", "frm", "gdb", "gtable", "kexi", "kexic", "kexis", "ldb", "mda", "mdb", "adp",
	"mde", "mdf", "myd", "myi", "ncf", "nsf", "ntf", "nv2", "odb", "ora", "pcontact", "pdb", "pdi", "pdx", "prc", "sql", "rec", "rel",
	"rin", "sdb", "sdf", "sqlite", "udl", "wadata", "waindx", "wamodel", "wajournal", "wdb", "wmdb",
}

var PUBLISHING_EXT = []string{
	"ai", "ave", "cdr", "chp", "pub", "sty", "cap", "cif", "vgr", "frm", "cpt", "dtp", "fm", "gdraw", "ildoc",
	"indd", "mcf", "pmd", "ppp", "psd", "pub", "qxd", "sla", "scd", "wfp", "wlmp", "wve", "xcf", "krita",
}

var DOCUMENTS_EXT = []string{
	"0", "1st", "600", "602", "abw", "acl", "afp", "ami", "amigaguide", "ans", "asc", "aww", "ccf", "csv", "cwk", "dbk", "dita", "doc", "docm", "docx", "dot",
	"dotx", "dwd", "egt", "epub", "ezw", "fdx", "ftm", "ftx", "gdoc", "hwp", "hwpml", "lwp", "mbp", "md", "me", "mcw", "mobi", "nb", "nb",
	"nbp", "neis", "odm	", "odoc", "odt", "osheet", "ott", "omm", "pages", "pap", "pdax", "pdf", "quox", "radix-64", "rtf", "rpt", "sdw", "se", "stw", "sxw",
	"tex", "info", "troff", "txt", "uof", "uoml", "via", "wpd", "wps", "wpt", "wrd", "wrf", "wri", "xhtml", "xml", "xps",
}

var FINANCIAL_DATA_EXT = []string{
	"myo", "myob", "tax", "ynab", "ifx", "ofx", "qfx", "qif",
}

var FONTS_EXT = []string{
	"abf", "afm", "bdf", "bmf", "fnt", "fon", "mgf", "otf", "pcf", "pfa", "pfb", "pfm",
	"afm", "fond", "sfd", "snf", "tdf", "tfm", "ttf", "ttc", "ufo", "woff",
}

var COLORS_EXT = []string{
	"act", "ase", "gpl", "pal", "icc", "icm",
}

var PICTURES_EXT = []string{
	"art", "blp", "bmp", "bti", "cd5", "cit", "cpt", "cr2", "clip", "cpl", "dds", "dib", "djvu", "egt", "exif", "grf", "icns", "ico", "iff",
	"ilbm", "lbm", "jng", "jpeg", "jfif", "jpg", "jp2", "jps", "lbm", "max", "miff", "mng", "msp", "nitf", "otb", "pbm", "pc1", "pc2", "pc3",
	"pcf", "pcx", "pdn", "pgm", "pi1", "pi2", "pi3", "pict", "pct", "png", "pnm", "pns", "ppm", "psb", "psp", "px", "pxm", "pxr", "qfx", "raw",
	"rle", "sct", "sgi", "rgb", "int", "bw", "tga", "targa", "icb", "vda", "vst", "pix", "tif", "tiff", "ep", "vtf", "xbm", "xcf", "xpm", "zif",
}

var VECTOR_GRAPHICS_EXT = []string{
	"3dv", "amf", "awg", "ai", "cgm", "cdr", "cmx", "dp", "dxf", "e2d", "egt", "eps", "fs", "gbr", "odg", "svg",
	"stl", "wrl", "x3d", "sxd", "v2d", "vdoc", "vsd", "vsdx", "vnd", "wmf", "emf", "art", "xar",
}

var _3_D_GRAPHICS_EXT = []string{
	"3dmf", "3dm", "3mf", "3ds", "abc", "ac", "amf", "an8", "aoi", "asm", "b3d", "blend", "block", "bmd3", "bdl4", "bdl", "brres", "bfres", "c4d", "cal3d",
	"ccp4", "cfl", "cob", "core3d", "ctm", "dae", "dff", "dpm", "dts", "egg", "fact", "fbx", "g", "glb", "glm", "gltf", "iob", "jas", "lwo", "lws", "lxf",
	"lxo", "ma", "max", "mb", "md2", "md3", "md5", "mdx", "mesh", "mesh", "mm3d", "mpo", "mrc", "nif", "obj", "off",
	"ogex", "ply", "prc", "prt", "pov", "r3d", "rwx", "sia", "sib", "skp", "sldasm", "sldprt", "smd", "u3d", "usd", "usda",
	"usdc", "usdz", "vim", "vrml97", "wrl", "vue", "vwx", "wings", "w3d", "x", "x3d", "z3d",
}

var SHORTCUTS_EXT = []string{
	"alias", "jnlp", "lnk", "appref-ms", "url", "webloc", "sym", "desktop",
}

var EXECUTABLE_FILES_EXT = []string{
	"8bf", "a", "out", "o", "so", "apk", "app", "bac", "bpl", "bundle", "class", "coff", "dcu", "dll", "dol", "ear", "elf",
	"expander", "exe", "ipa", "jeff", "jar", "xpi", "dylib", "bundle", "nlm", "obb", "rll", "s1es", "vap", "war", "xbe", "xap",
	"xcoff", "xex", "vbx", "ocx", "tlb",
}

var PRESENTATION_EXT = []string{
	"gslides", "key", "nb", "nbp", "odp", "otp", "pez", "pot", "pps", "ppt", "pptx", "prz",
	"sdd", "shf", "show", "shw", "slp", "sspss", "sti", "sxi", "thmx", "watch",
}

var AUDIO_EXT = []string{
	"8svx", "16svx", "aiff", "aif", "aifc", "au", "bwf", "cdda", "raw", "wav", "ra", "rm", "flac", "la", "pac", "ape", "ofr", "ofs", "off",
	"rka", "shn", "tak", "thd", "tta", "wv", "wma", "brstm", "dts", "dtshd", "dtsma", "ast", "aw", "psf", "ac3", "amr", "mp1", "mp2", "mp3",
	"spx", "gsm", "wma", "aac", "mpc", "vqf", "ots", "swa", "vox", "voc", "dwd", "smp", "ogg",
}

var PLAYLIST_EXT = []string{
	"aimppl", "asx", "ram", "xpl", "xspf", "zpl", "m3u", "pls",
}

var SOURCE_CODE_EXT = []string{
	"ada", "adb", "ads", "asm", "s", "bas", "bb", "bmx", "c", "clj", "cls", "cob", "cbl", "cpp", "cc", "cxx", "c", "cbp", "cs", "csproj", "d", "dba",
	"dbpro123", "e", "efs", "egt", "el", "for", "ftn", "f", "f77", "f90", "frm", "frx", "fth", "ged", "gm6", "gmd", "gmk", "gml", "go", "h", "hpp", "hxx",
	"hs", "i", "inc", "java", "l", "lgt", "lisp", "m", "m4", "ml", "msqr", "n", "nb", "p", "pas", "pp", "php", "pisrc", "piv", "pl",
	"pm", "pli", "pl1", "prg", "pro", "pol", "py", "r", "red", "reds", "rb", "resx", "rc", "rkt", "rktl", "scala", "sci", "sce", "scm", "sd7", "skb",
	"skc", "skd", "skf", "skg", "ski", "skk", "skm", "sko", "skp", "skq", "sks", "skt", "skz", "sln", "spin", "stk", "swg", "tcl", "vap", "vb", "vbg", "vbp",
	"vip", "vbproj", "vcproj", "vdproj", "xpl", "xq", "xsl", "y", "js", "jsx", "sass", "less",
}

var SPREADSHEETS_EXT = []string{
	"123", "ab2", "ab3", "aws", "bcsv", "clf", "cell", "csv", "gsheet", "numbers", "gnumeric", "lcw", "ods", "ots",
	"qpw", "sdc", "slk", "stc", "sxc", "tab", "vc", "wk1", "wk3", "wk4", "wks", "wks", "wq1", "xlk", "xls", "xlsb", "xlsm", "xlsx",
	"xlr", "xlt", "xltm", "xlw", "tsv",
}

var VIDEOS_EXT = []string{
	"aaf", "3gp", "gif", "asf", "avchd", "avi", "bik", "cam", "collab", "dat", "dsh", "dvr-ms", "flv", "m1v", "m2v", "fla", "flr", "sol",
	"m4v", "mkv", "wrap", "mng", "mov", "mpeg", "mpg", "mpe", "thp", "mxf", "roq", "nsv", "ogg", "rm", "svi", "smi", "smk", "swf", "wmv",
	"wtv", "yuv", "webm", "mp4",
}

var VIRTUAL_MACHINES_EXT = []string{
	"vfd", "vhd", "vud", "vmc", "vsv", "vmdk", "dsk", "nvram", "vmem", "vmsd", "vmsn", "vmss",
	"std", "vmtm", "vmx", "vmxf", "vdi", "vbox-extpack", "hdd", "pvs", "sav", "cow", "qcow", "qcow2", "qed",
}

var WEB_PAGES_EXT = []string{
	"dtd", "html", "htm", "xhtml", "xht", "mht", "mhtml", ".maff", "asp", "aspx", "adp", "bml",
	"cfm", "cgi", "ihtml", "jsp", "lasso", "las", "lassoapp", "pl", "rna", "shtml", "stm",
}
