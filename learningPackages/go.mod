module learningPackages

// The line module learnpackage specifies that the module’s name is learnpackage.
// learnpackage will be the base path to import any package created inside this module.
go 1.23.2

replace level_0 => ../level_0

require level_0 v0.0.0-00010101000000-000000000000 // indirect
