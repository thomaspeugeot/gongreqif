// generated code - do not edit
package models

import (
	"bufio"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_ALTERNATIVEID = make(map[string]*ALTERNATIVEID)
var __gong__map_ATTRIBUTEDEFINITIONBOOLEAN = make(map[string]*ATTRIBUTEDEFINITIONBOOLEAN)
var __gong__map_ATTRIBUTEDEFINITIONDATE = make(map[string]*ATTRIBUTEDEFINITIONDATE)
var __gong__map_ATTRIBUTEDEFINITIONENUMERATION = make(map[string]*ATTRIBUTEDEFINITIONENUMERATION)
var __gong__map_ATTRIBUTEDEFINITIONINTEGER = make(map[string]*ATTRIBUTEDEFINITIONINTEGER)
var __gong__map_ATTRIBUTEDEFINITIONREAL = make(map[string]*ATTRIBUTEDEFINITIONREAL)
var __gong__map_ATTRIBUTEDEFINITIONSTRING = make(map[string]*ATTRIBUTEDEFINITIONSTRING)
var __gong__map_ATTRIBUTEDEFINITIONXHTML = make(map[string]*ATTRIBUTEDEFINITIONXHTML)
var __gong__map_ATTRIBUTEVALUEBOOLEAN = make(map[string]*ATTRIBUTEVALUEBOOLEAN)
var __gong__map_ATTRIBUTEVALUEDATE = make(map[string]*ATTRIBUTEVALUEDATE)
var __gong__map_ATTRIBUTEVALUEENUMERATION = make(map[string]*ATTRIBUTEVALUEENUMERATION)
var __gong__map_ATTRIBUTEVALUEINTEGER = make(map[string]*ATTRIBUTEVALUEINTEGER)
var __gong__map_ATTRIBUTEVALUEREAL = make(map[string]*ATTRIBUTEVALUEREAL)
var __gong__map_ATTRIBUTEVALUESTRING = make(map[string]*ATTRIBUTEVALUESTRING)
var __gong__map_ATTRIBUTEVALUEXHTML = make(map[string]*ATTRIBUTEVALUEXHTML)
var __gong__map_CHILDREN = make(map[string]*CHILDREN)
var __gong__map_CORECONTENT = make(map[string]*CORECONTENT)
var __gong__map_DATATYPEDEFINITIONBOOLEAN = make(map[string]*DATATYPEDEFINITIONBOOLEAN)
var __gong__map_DATATYPEDEFINITIONDATE = make(map[string]*DATATYPEDEFINITIONDATE)
var __gong__map_DATATYPEDEFINITIONENUMERATION = make(map[string]*DATATYPEDEFINITIONENUMERATION)
var __gong__map_DATATYPEDEFINITIONINTEGER = make(map[string]*DATATYPEDEFINITIONINTEGER)
var __gong__map_DATATYPEDEFINITIONREAL = make(map[string]*DATATYPEDEFINITIONREAL)
var __gong__map_DATATYPEDEFINITIONSTRING = make(map[string]*DATATYPEDEFINITIONSTRING)
var __gong__map_DATATYPEDEFINITIONXHTML = make(map[string]*DATATYPEDEFINITIONXHTML)
var __gong__map_DATATYPES = make(map[string]*DATATYPES)
var __gong__map_DEFAULTVALUE = make(map[string]*DEFAULTVALUE)
var __gong__map_DEFINITION = make(map[string]*DEFINITION)
var __gong__map_EDITABLEATTS = make(map[string]*EDITABLEATTS)
var __gong__map_EMBEDDEDVALUE = make(map[string]*EMBEDDEDVALUE)
var __gong__map_ENUMVALUE = make(map[string]*ENUMVALUE)
var __gong__map_OBJECT = make(map[string]*OBJECT)
var __gong__map_PROPERTIES = make(map[string]*PROPERTIES)
var __gong__map_RELATIONGROUP = make(map[string]*RELATIONGROUP)
var __gong__map_RELATIONGROUPTYPE = make(map[string]*RELATIONGROUPTYPE)
var __gong__map_REQIF = make(map[string]*REQIF)
var __gong__map_REQIFCONTENT = make(map[string]*REQIFCONTENT)
var __gong__map_REQIFHEADER = make(map[string]*REQIFHEADER)
var __gong__map_REQIFTOOLEXTENSION = make(map[string]*REQIFTOOLEXTENSION)
var __gong__map_REQTYPE = make(map[string]*REQTYPE)
var __gong__map_SOURCE = make(map[string]*SOURCE)
var __gong__map_SOURCESPECIFICATION = make(map[string]*SOURCESPECIFICATION)
var __gong__map_SPECATTRIBUTES = make(map[string]*SPECATTRIBUTES)
var __gong__map_SPECHIERARCHY = make(map[string]*SPECHIERARCHY)
var __gong__map_SPECIFICATION = make(map[string]*SPECIFICATION)
var __gong__map_SPECIFICATIONS = make(map[string]*SPECIFICATIONS)
var __gong__map_SPECIFICATIONTYPE = make(map[string]*SPECIFICATIONTYPE)
var __gong__map_SPECIFIEDVALUES = make(map[string]*SPECIFIEDVALUES)
var __gong__map_SPECOBJECT = make(map[string]*SPECOBJECT)
var __gong__map_SPECOBJECTS = make(map[string]*SPECOBJECTS)
var __gong__map_SPECOBJECTTYPE = make(map[string]*SPECOBJECTTYPE)
var __gong__map_SPECRELATION = make(map[string]*SPECRELATION)
var __gong__map_SPECRELATIONGROUPS = make(map[string]*SPECRELATIONGROUPS)
var __gong__map_SPECRELATIONS = make(map[string]*SPECRELATIONS)
var __gong__map_SPECRELATIONTYPE = make(map[string]*SPECRELATIONTYPE)
var __gong__map_SPECTYPES = make(map[string]*SPECTYPES)
var __gong__map_TARGET = make(map[string]*TARGET)
var __gong__map_TARGETSPECIFICATION = make(map[string]*TARGETSPECIFICATION)
var __gong__map_THEHEADER = make(map[string]*THEHEADER)
var __gong__map_TOOLEXTENSIONS = make(map[string]*TOOLEXTENSIONS)
var __gong__map_VALUES = make(map[string]*VALUES)
var __gong__map_XHTMLCONTENT = make(map[string]*XHTMLCONTENT)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "ALTERNATIVEID":
										instanceALTERNATIVEID := (&ALTERNATIVEID{Name: instanceName}).Stage(stage)
										instance = any(instanceALTERNATIVEID)
										__gong__map_ALTERNATIVEID[identifier] = instanceALTERNATIVEID
									case "ATTRIBUTEDEFINITIONBOOLEAN":
										instanceATTRIBUTEDEFINITIONBOOLEAN := (&ATTRIBUTEDEFINITIONBOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONBOOLEAN)
										__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier] = instanceATTRIBUTEDEFINITIONBOOLEAN
									case "ATTRIBUTEDEFINITIONDATE":
										instanceATTRIBUTEDEFINITIONDATE := (&ATTRIBUTEDEFINITIONDATE{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONDATE)
										__gong__map_ATTRIBUTEDEFINITIONDATE[identifier] = instanceATTRIBUTEDEFINITIONDATE
									case "ATTRIBUTEDEFINITIONENUMERATION":
										instanceATTRIBUTEDEFINITIONENUMERATION := (&ATTRIBUTEDEFINITIONENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONENUMERATION)
										__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier] = instanceATTRIBUTEDEFINITIONENUMERATION
									case "ATTRIBUTEDEFINITIONINTEGER":
										instanceATTRIBUTEDEFINITIONINTEGER := (&ATTRIBUTEDEFINITIONINTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONINTEGER)
										__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier] = instanceATTRIBUTEDEFINITIONINTEGER
									case "ATTRIBUTEDEFINITIONREAL":
										instanceATTRIBUTEDEFINITIONREAL := (&ATTRIBUTEDEFINITIONREAL{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONREAL)
										__gong__map_ATTRIBUTEDEFINITIONREAL[identifier] = instanceATTRIBUTEDEFINITIONREAL
									case "ATTRIBUTEDEFINITIONSTRING":
										instanceATTRIBUTEDEFINITIONSTRING := (&ATTRIBUTEDEFINITIONSTRING{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONSTRING)
										__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier] = instanceATTRIBUTEDEFINITIONSTRING
									case "ATTRIBUTEDEFINITIONXHTML":
										instanceATTRIBUTEDEFINITIONXHTML := (&ATTRIBUTEDEFINITIONXHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEDEFINITIONXHTML)
										__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier] = instanceATTRIBUTEDEFINITIONXHTML
									case "ATTRIBUTEVALUEBOOLEAN":
										instanceATTRIBUTEVALUEBOOLEAN := (&ATTRIBUTEVALUEBOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEBOOLEAN)
										__gong__map_ATTRIBUTEVALUEBOOLEAN[identifier] = instanceATTRIBUTEVALUEBOOLEAN
									case "ATTRIBUTEVALUEDATE":
										instanceATTRIBUTEVALUEDATE := (&ATTRIBUTEVALUEDATE{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEDATE)
										__gong__map_ATTRIBUTEVALUEDATE[identifier] = instanceATTRIBUTEVALUEDATE
									case "ATTRIBUTEVALUEENUMERATION":
										instanceATTRIBUTEVALUEENUMERATION := (&ATTRIBUTEVALUEENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEENUMERATION)
										__gong__map_ATTRIBUTEVALUEENUMERATION[identifier] = instanceATTRIBUTEVALUEENUMERATION
									case "ATTRIBUTEVALUEINTEGER":
										instanceATTRIBUTEVALUEINTEGER := (&ATTRIBUTEVALUEINTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEINTEGER)
										__gong__map_ATTRIBUTEVALUEINTEGER[identifier] = instanceATTRIBUTEVALUEINTEGER
									case "ATTRIBUTEVALUEREAL":
										instanceATTRIBUTEVALUEREAL := (&ATTRIBUTEVALUEREAL{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEREAL)
										__gong__map_ATTRIBUTEVALUEREAL[identifier] = instanceATTRIBUTEVALUEREAL
									case "ATTRIBUTEVALUESTRING":
										instanceATTRIBUTEVALUESTRING := (&ATTRIBUTEVALUESTRING{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUESTRING)
										__gong__map_ATTRIBUTEVALUESTRING[identifier] = instanceATTRIBUTEVALUESTRING
									case "ATTRIBUTEVALUEXHTML":
										instanceATTRIBUTEVALUEXHTML := (&ATTRIBUTEVALUEXHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTEVALUEXHTML)
										__gong__map_ATTRIBUTEVALUEXHTML[identifier] = instanceATTRIBUTEVALUEXHTML
									case "CHILDREN":
										instanceCHILDREN := (&CHILDREN{Name: instanceName}).Stage(stage)
										instance = any(instanceCHILDREN)
										__gong__map_CHILDREN[identifier] = instanceCHILDREN
									case "CORECONTENT":
										instanceCORECONTENT := (&CORECONTENT{Name: instanceName}).Stage(stage)
										instance = any(instanceCORECONTENT)
										__gong__map_CORECONTENT[identifier] = instanceCORECONTENT
									case "DATATYPEDEFINITIONBOOLEAN":
										instanceDATATYPEDEFINITIONBOOLEAN := (&DATATYPEDEFINITIONBOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONBOOLEAN)
										__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier] = instanceDATATYPEDEFINITIONBOOLEAN
									case "DATATYPEDEFINITIONDATE":
										instanceDATATYPEDEFINITIONDATE := (&DATATYPEDEFINITIONDATE{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONDATE)
										__gong__map_DATATYPEDEFINITIONDATE[identifier] = instanceDATATYPEDEFINITIONDATE
									case "DATATYPEDEFINITIONENUMERATION":
										instanceDATATYPEDEFINITIONENUMERATION := (&DATATYPEDEFINITIONENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONENUMERATION)
										__gong__map_DATATYPEDEFINITIONENUMERATION[identifier] = instanceDATATYPEDEFINITIONENUMERATION
									case "DATATYPEDEFINITIONINTEGER":
										instanceDATATYPEDEFINITIONINTEGER := (&DATATYPEDEFINITIONINTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONINTEGER)
										__gong__map_DATATYPEDEFINITIONINTEGER[identifier] = instanceDATATYPEDEFINITIONINTEGER
									case "DATATYPEDEFINITIONREAL":
										instanceDATATYPEDEFINITIONREAL := (&DATATYPEDEFINITIONREAL{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONREAL)
										__gong__map_DATATYPEDEFINITIONREAL[identifier] = instanceDATATYPEDEFINITIONREAL
									case "DATATYPEDEFINITIONSTRING":
										instanceDATATYPEDEFINITIONSTRING := (&DATATYPEDEFINITIONSTRING{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONSTRING)
										__gong__map_DATATYPEDEFINITIONSTRING[identifier] = instanceDATATYPEDEFINITIONSTRING
									case "DATATYPEDEFINITIONXHTML":
										instanceDATATYPEDEFINITIONXHTML := (&DATATYPEDEFINITIONXHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPEDEFINITIONXHTML)
										__gong__map_DATATYPEDEFINITIONXHTML[identifier] = instanceDATATYPEDEFINITIONXHTML
									case "DATATYPES":
										instanceDATATYPES := (&DATATYPES{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPES)
										__gong__map_DATATYPES[identifier] = instanceDATATYPES
									case "DEFAULTVALUE":
										instanceDEFAULTVALUE := (&DEFAULTVALUE{Name: instanceName}).Stage(stage)
										instance = any(instanceDEFAULTVALUE)
										__gong__map_DEFAULTVALUE[identifier] = instanceDEFAULTVALUE
									case "DEFINITION":
										instanceDEFINITION := (&DEFINITION{Name: instanceName}).Stage(stage)
										instance = any(instanceDEFINITION)
										__gong__map_DEFINITION[identifier] = instanceDEFINITION
									case "EDITABLEATTS":
										instanceEDITABLEATTS := (&EDITABLEATTS{Name: instanceName}).Stage(stage)
										instance = any(instanceEDITABLEATTS)
										__gong__map_EDITABLEATTS[identifier] = instanceEDITABLEATTS
									case "EMBEDDEDVALUE":
										instanceEMBEDDEDVALUE := (&EMBEDDEDVALUE{Name: instanceName}).Stage(stage)
										instance = any(instanceEMBEDDEDVALUE)
										__gong__map_EMBEDDEDVALUE[identifier] = instanceEMBEDDEDVALUE
									case "ENUMVALUE":
										instanceENUMVALUE := (&ENUMVALUE{Name: instanceName}).Stage(stage)
										instance = any(instanceENUMVALUE)
										__gong__map_ENUMVALUE[identifier] = instanceENUMVALUE
									case "OBJECT":
										instanceOBJECT := (&OBJECT{Name: instanceName}).Stage(stage)
										instance = any(instanceOBJECT)
										__gong__map_OBJECT[identifier] = instanceOBJECT
									case "PROPERTIES":
										instancePROPERTIES := (&PROPERTIES{Name: instanceName}).Stage(stage)
										instance = any(instancePROPERTIES)
										__gong__map_PROPERTIES[identifier] = instancePROPERTIES
									case "RELATIONGROUP":
										instanceRELATIONGROUP := (&RELATIONGROUP{Name: instanceName}).Stage(stage)
										instance = any(instanceRELATIONGROUP)
										__gong__map_RELATIONGROUP[identifier] = instanceRELATIONGROUP
									case "RELATIONGROUPTYPE":
										instanceRELATIONGROUPTYPE := (&RELATIONGROUPTYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceRELATIONGROUPTYPE)
										__gong__map_RELATIONGROUPTYPE[identifier] = instanceRELATIONGROUPTYPE
									case "REQIF":
										instanceREQIF := (&REQIF{Name: instanceName}).Stage(stage)
										instance = any(instanceREQIF)
										__gong__map_REQIF[identifier] = instanceREQIF
									case "REQIFCONTENT":
										instanceREQIFCONTENT := (&REQIFCONTENT{Name: instanceName}).Stage(stage)
										instance = any(instanceREQIFCONTENT)
										__gong__map_REQIFCONTENT[identifier] = instanceREQIFCONTENT
									case "REQIFHEADER":
										instanceREQIFHEADER := (&REQIFHEADER{Name: instanceName}).Stage(stage)
										instance = any(instanceREQIFHEADER)
										__gong__map_REQIFHEADER[identifier] = instanceREQIFHEADER
									case "REQIFTOOLEXTENSION":
										instanceREQIFTOOLEXTENSION := (&REQIFTOOLEXTENSION{Name: instanceName}).Stage(stage)
										instance = any(instanceREQIFTOOLEXTENSION)
										__gong__map_REQIFTOOLEXTENSION[identifier] = instanceREQIFTOOLEXTENSION
									case "REQTYPE":
										instanceREQTYPE := (&REQTYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceREQTYPE)
										__gong__map_REQTYPE[identifier] = instanceREQTYPE
									case "SOURCE":
										instanceSOURCE := (&SOURCE{Name: instanceName}).Stage(stage)
										instance = any(instanceSOURCE)
										__gong__map_SOURCE[identifier] = instanceSOURCE
									case "SOURCESPECIFICATION":
										instanceSOURCESPECIFICATION := (&SOURCESPECIFICATION{Name: instanceName}).Stage(stage)
										instance = any(instanceSOURCESPECIFICATION)
										__gong__map_SOURCESPECIFICATION[identifier] = instanceSOURCESPECIFICATION
									case "SPECATTRIBUTES":
										instanceSPECATTRIBUTES := (&SPECATTRIBUTES{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECATTRIBUTES)
										__gong__map_SPECATTRIBUTES[identifier] = instanceSPECATTRIBUTES
									case "SPECHIERARCHY":
										instanceSPECHIERARCHY := (&SPECHIERARCHY{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECHIERARCHY)
										__gong__map_SPECHIERARCHY[identifier] = instanceSPECHIERARCHY
									case "SPECIFICATION":
										instanceSPECIFICATION := (&SPECIFICATION{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFICATION)
										__gong__map_SPECIFICATION[identifier] = instanceSPECIFICATION
									case "SPECIFICATIONS":
										instanceSPECIFICATIONS := (&SPECIFICATIONS{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFICATIONS)
										__gong__map_SPECIFICATIONS[identifier] = instanceSPECIFICATIONS
									case "SPECIFICATIONTYPE":
										instanceSPECIFICATIONTYPE := (&SPECIFICATIONTYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFICATIONTYPE)
										__gong__map_SPECIFICATIONTYPE[identifier] = instanceSPECIFICATIONTYPE
									case "SPECIFIEDVALUES":
										instanceSPECIFIEDVALUES := (&SPECIFIEDVALUES{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFIEDVALUES)
										__gong__map_SPECIFIEDVALUES[identifier] = instanceSPECIFIEDVALUES
									case "SPECOBJECT":
										instanceSPECOBJECT := (&SPECOBJECT{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECOBJECT)
										__gong__map_SPECOBJECT[identifier] = instanceSPECOBJECT
									case "SPECOBJECTS":
										instanceSPECOBJECTS := (&SPECOBJECTS{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECOBJECTS)
										__gong__map_SPECOBJECTS[identifier] = instanceSPECOBJECTS
									case "SPECOBJECTTYPE":
										instanceSPECOBJECTTYPE := (&SPECOBJECTTYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECOBJECTTYPE)
										__gong__map_SPECOBJECTTYPE[identifier] = instanceSPECOBJECTTYPE
									case "SPECRELATION":
										instanceSPECRELATION := (&SPECRELATION{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECRELATION)
										__gong__map_SPECRELATION[identifier] = instanceSPECRELATION
									case "SPECRELATIONGROUPS":
										instanceSPECRELATIONGROUPS := (&SPECRELATIONGROUPS{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECRELATIONGROUPS)
										__gong__map_SPECRELATIONGROUPS[identifier] = instanceSPECRELATIONGROUPS
									case "SPECRELATIONS":
										instanceSPECRELATIONS := (&SPECRELATIONS{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECRELATIONS)
										__gong__map_SPECRELATIONS[identifier] = instanceSPECRELATIONS
									case "SPECRELATIONTYPE":
										instanceSPECRELATIONTYPE := (&SPECRELATIONTYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECRELATIONTYPE)
										__gong__map_SPECRELATIONTYPE[identifier] = instanceSPECRELATIONTYPE
									case "SPECTYPES":
										instanceSPECTYPES := (&SPECTYPES{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECTYPES)
										__gong__map_SPECTYPES[identifier] = instanceSPECTYPES
									case "TARGET":
										instanceTARGET := (&TARGET{Name: instanceName}).Stage(stage)
										instance = any(instanceTARGET)
										__gong__map_TARGET[identifier] = instanceTARGET
									case "TARGETSPECIFICATION":
										instanceTARGETSPECIFICATION := (&TARGETSPECIFICATION{Name: instanceName}).Stage(stage)
										instance = any(instanceTARGETSPECIFICATION)
										__gong__map_TARGETSPECIFICATION[identifier] = instanceTARGETSPECIFICATION
									case "THEHEADER":
										instanceTHEHEADER := (&THEHEADER{Name: instanceName}).Stage(stage)
										instance = any(instanceTHEHEADER)
										__gong__map_THEHEADER[identifier] = instanceTHEHEADER
									case "TOOLEXTENSIONS":
										instanceTOOLEXTENSIONS := (&TOOLEXTENSIONS{Name: instanceName}).Stage(stage)
										instance = any(instanceTOOLEXTENSIONS)
										__gong__map_TOOLEXTENSIONS[identifier] = instanceTOOLEXTENSIONS
									case "VALUES":
										instanceVALUES := (&VALUES{Name: instanceName}).Stage(stage)
										instance = any(instanceVALUES)
										__gong__map_VALUES[identifier] = instanceVALUES
									case "XHTMLCONTENT":
										instanceXHTMLCONTENT := (&XHTMLCONTENT{Name: instanceName}).Stage(stage)
										instance = any(instanceXHTMLCONTENT)
										__gong__map_XHTMLCONTENT[identifier] = instanceXHTMLCONTENT
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "ALTERNATIVEID":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONBOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONDATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONINTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONREAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONSTRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEDEFINITIONXHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEBOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEDATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEINTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEREAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUESTRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTEVALUEXHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CHILDREN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CORECONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONBOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONDATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONINTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONREAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONSTRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPEDEFINITIONXHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DEFAULTVALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DEFINITION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EDITABLEATTS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EMBEDDEDVALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ENUMVALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "OBJECT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "PROPERTIES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATIONGROUP":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATIONGROUPTYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQIF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQIFCONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQIFHEADER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQIFTOOLEXTENSION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQTYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SOURCE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SOURCESPECIFICATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECATTRIBUTES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECHIERARCHY":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATIONTYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFIEDVALUES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECOBJECT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECOBJECTS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECOBJECTTYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECRELATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECRELATIONGROUPS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECRELATIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECRELATIONTYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECTYPES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TARGET":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TARGETSPECIFICATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "THEHEADER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TOOLEXTENSIONS":
							switch fieldName {
							// insertion point for date assign code
							}
						case "VALUES":
							switch fieldName {
							// insertion point for date assign code
							}
						case "XHTMLCONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "ALTERNATIVEID":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONBOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONDATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONINTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONREAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONSTRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEDEFINITIONXHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEBOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEDATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEINTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEREAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUESTRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTEVALUEXHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CHILDREN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPECHIERARCHY":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECHIERARCHY[targetIdentifier]
							__gong__map_CHILDREN[identifier].SPECHIERARCHY =
								append(__gong__map_CHILDREN[identifier].SPECHIERARCHY, target)
						}
					case "CORECONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONBOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONDATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONINTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONREAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONSTRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPEDEFINITIONXHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "DATATYPEDEFINITIONBOOLEAN":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONBOOLEAN[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONBOOLEAN =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONBOOLEAN, target)
						case "DATATYPEDEFINITIONDATE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONDATE[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONDATE =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONDATE, target)
						case "DATATYPEDEFINITIONENUMERATION":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONENUMERATION[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONENUMERATION =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONENUMERATION, target)
						case "DATATYPEDEFINITIONINTEGER":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONINTEGER[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONINTEGER =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONINTEGER, target)
						case "DATATYPEDEFINITIONREAL":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONREAL[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONREAL =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONREAL, target)
						case "DATATYPEDEFINITIONSTRING":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONSTRING[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONSTRING =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONSTRING, target)
						case "DATATYPEDEFINITIONXHTML":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DATATYPEDEFINITIONXHTML[targetIdentifier]
							__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONXHTML =
								append(__gong__map_DATATYPES[identifier].DATATYPEDEFINITIONXHTML, target)
						}
					case "DEFAULTVALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DEFINITION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EDITABLEATTS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EMBEDDEDVALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ENUMVALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "OBJECT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "PROPERTIES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATIONGROUP":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATIONGROUPTYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQIF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQIFCONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQIFHEADER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQIFTOOLEXTENSION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQTYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SOURCE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SOURCESPECIFICATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECATTRIBUTES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ATTRIBUTEDEFINITIONBOOLEAN":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONBOOLEAN[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONBOOLEAN =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONBOOLEAN, target)
						case "ATTRIBUTEDEFINITIONDATE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONDATE[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONDATE =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONDATE, target)
						case "ATTRIBUTEDEFINITIONENUMERATION":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONENUMERATION[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONENUMERATION =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONENUMERATION, target)
						case "ATTRIBUTEDEFINITIONINTEGER":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONINTEGER[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONINTEGER =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONINTEGER, target)
						case "ATTRIBUTEDEFINITIONREAL":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONREAL[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONREAL =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONREAL, target)
						case "ATTRIBUTEDEFINITIONSTRING":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONSTRING[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONSTRING =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONSTRING, target)
						case "ATTRIBUTEDEFINITIONXHTML":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ATTRIBUTEDEFINITIONXHTML[targetIdentifier]
							__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONXHTML =
								append(__gong__map_SPECATTRIBUTES[identifier].ATTRIBUTEDEFINITIONXHTML, target)
						}
					case "SPECHIERARCHY":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFICATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFICATIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPECIFICATION":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECIFICATION[targetIdentifier]
							__gong__map_SPECIFICATIONS[identifier].SPECIFICATION =
								append(__gong__map_SPECIFICATIONS[identifier].SPECIFICATION, target)
						}
					case "SPECIFICATIONTYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFIEDVALUES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ENUMVALUE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ENUMVALUE[targetIdentifier]
							__gong__map_SPECIFIEDVALUES[identifier].ENUMVALUE =
								append(__gong__map_SPECIFIEDVALUES[identifier].ENUMVALUE, target)
						}
					case "SPECOBJECT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECOBJECTS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SPECOBJECT":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECOBJECT[targetIdentifier]
							__gong__map_SPECOBJECTS[identifier].SPECOBJECT =
								append(__gong__map_SPECOBJECTS[identifier].SPECOBJECT, target)
						}
					case "SPECOBJECTTYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECRELATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECRELATIONGROUPS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RELATIONGROUP":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_RELATIONGROUP[targetIdentifier]
							__gong__map_SPECRELATIONGROUPS[identifier].RELATIONGROUP =
								append(__gong__map_SPECRELATIONGROUPS[identifier].RELATIONGROUP, target)
						}
					case "SPECRELATIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECRELATIONTYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECTYPES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RELATIONGROUPTYPE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_RELATIONGROUPTYPE[targetIdentifier]
							__gong__map_SPECTYPES[identifier].RELATIONGROUPTYPE =
								append(__gong__map_SPECTYPES[identifier].RELATIONGROUPTYPE, target)
						case "SPECOBJECTTYPE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECOBJECTTYPE[targetIdentifier]
							__gong__map_SPECTYPES[identifier].SPECOBJECTTYPE =
								append(__gong__map_SPECTYPES[identifier].SPECOBJECTTYPE, target)
						case "SPECRELATIONTYPE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECRELATIONTYPE[targetIdentifier]
							__gong__map_SPECTYPES[identifier].SPECRELATIONTYPE =
								append(__gong__map_SPECTYPES[identifier].SPECRELATIONTYPE, target)
						case "SPECIFICATIONTYPE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SPECIFICATIONTYPE[targetIdentifier]
							__gong__map_SPECTYPES[identifier].SPECIFICATIONTYPE =
								append(__gong__map_SPECTYPES[identifier].SPECIFICATIONTYPE, target)
						}
					case "TARGET":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TARGETSPECIFICATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "THEHEADER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TOOLEXTENSIONS":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "REQIFTOOLEXTENSION":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_REQIFTOOLEXTENSION[targetIdentifier]
							__gong__map_TOOLEXTENSIONS[identifier].REQIFTOOLEXTENSION =
								append(__gong__map_TOOLEXTENSIONS[identifier].REQIFTOOLEXTENSION, target)
						}
					case "VALUES":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "XHTMLCONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "ALTERNATIVEID":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVEID[identifier].Name = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVEID[identifier].IDENTIFIERAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONSTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEDEFINITIONXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].LONGNAMEAttr = fielValue
				}
			case "ATTRIBUTEVALUEBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEBOOLEAN[identifier].Name = fielValue
				}
			case "ATTRIBUTEVALUEDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEDATE[identifier].Name = fielValue
				case "THEVALUEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEDATE[identifier].THEVALUEAttr = fielValue
				}
			case "ATTRIBUTEVALUEENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEENUMERATION[identifier].Name = fielValue
				}
			case "ATTRIBUTEVALUEINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEINTEGER[identifier].Name = fielValue
				case "THEVALUEAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEVALUEINTEGER[identifier].THEVALUEAttr = int(exprSign) * int(fielValue)
				}
			case "ATTRIBUTEVALUEREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEREAL[identifier].Name = fielValue
				case "THEVALUEAttr":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEVALUEREAL[identifier].THEVALUEAttr = exprSign * fielValue
				}
			case "ATTRIBUTEVALUESTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUESTRING[identifier].Name = fielValue
				case "THEVALUEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUESTRING[identifier].THEVALUEAttr = fielValue
				}
			case "ATTRIBUTEVALUEXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTEVALUEXHTML[identifier].Name = fielValue
				}
			case "CHILDREN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CHILDREN[identifier].Name = fielValue
				}
			case "CORECONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CORECONTENT[identifier].Name = fielValue
				}
			case "DATATYPEDEFINITIONBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].LONGNAMEAttr = fielValue
				}
			case "DATATYPEDEFINITIONDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONDATE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONDATE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONDATE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONDATE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONDATE[identifier].LONGNAMEAttr = fielValue
				}
			case "DATATYPEDEFINITIONENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].LONGNAMEAttr = fielValue
				}
			case "DATATYPEDEFINITIONINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].LONGNAMEAttr = fielValue
				case "MAXAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].MAXAttr = int(exprSign) * int(fielValue)
				case "MINAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].MINAttr = int(exprSign) * int(fielValue)
				}
			case "DATATYPEDEFINITIONREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONREAL[identifier].Name = fielValue
				case "ACCURACYAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONREAL[identifier].ACCURACYAttr = int(exprSign) * int(fielValue)
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONREAL[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONREAL[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONREAL[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONREAL[identifier].LONGNAMEAttr = fielValue
				case "MAXAttr":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONREAL[identifier].MAXAttr = exprSign * fielValue
				case "MINAttr":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONREAL[identifier].MINAttr = exprSign * fielValue
				}
			case "DATATYPEDEFINITIONSTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].LONGNAMEAttr = fielValue
				case "MAXLENGTHAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].MAXLENGTHAttr = int(exprSign) * int(fielValue)
				}
			case "DATATYPEDEFINITIONXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].LONGNAMEAttr = fielValue
				}
			case "DATATYPES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPES[identifier].Name = fielValue
				}
			case "DEFAULTVALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DEFAULTVALUE[identifier].Name = fielValue
				}
			case "DEFINITION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DEFINITION[identifier].Name = fielValue
				case "ATTRIBUTEDEFINITIONBOOLEANREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DEFINITION[identifier].ATTRIBUTEDEFINITIONBOOLEANREF = fielValue
				}
			case "EDITABLEATTS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EDITABLEATTS[identifier].Name = fielValue
				}
			case "EMBEDDEDVALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDEDVALUE[identifier].Name = fielValue
				case "KEYAttr":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_EMBEDDEDVALUE[identifier].KEYAttr = int(exprSign) * int(fielValue)
				case "OTHERCONTENTAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDEDVALUE[identifier].OTHERCONTENTAttr = fielValue
				}
			case "ENUMVALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUMVALUE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUMVALUE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUMVALUE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUMVALUE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUMVALUE[identifier].LONGNAMEAttr = fielValue
				}
			case "OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_OBJECT[identifier].Name = fielValue
				case "SPECOBJECTREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_OBJECT[identifier].SPECOBJECTREF = fielValue
				}
			case "PROPERTIES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_PROPERTIES[identifier].Name = fielValue
				}
			case "RELATIONGROUP":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUP[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUP[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUP[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUP[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUP[identifier].LONGNAMEAttr = fielValue
				}
			case "RELATIONGROUPTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUPTYPE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUPTYPE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUPTYPE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUPTYPE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATIONGROUPTYPE[identifier].LONGNAMEAttr = fielValue
				}
			case "REQIF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIF[identifier].Name = fielValue
				}
			case "REQIFCONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFCONTENT[identifier].Name = fielValue
				}
			case "REQIFHEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].Name = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].IDENTIFIERAttr = fielValue
				case "COMMENT":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].COMMENT = fielValue
				case "CREATIONTIME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].CREATIONTIME = fielValue
				case "REPOSITORYID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].REPOSITORYID = fielValue
				case "REQIFTOOLID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].REQIFTOOLID = fielValue
				case "REQIFVERSION":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].REQIFVERSION = fielValue
				case "SOURCETOOLID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].SOURCETOOLID = fielValue
				case "TITLE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFHEADER[identifier].TITLE = fielValue
				}
			case "REQIFTOOLEXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQIFTOOLEXTENSION[identifier].Name = fielValue
				}
			case "REQTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQTYPE[identifier].Name = fielValue
				case "DATATYPEDEFINITIONBOOLEANREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQTYPE[identifier].DATATYPEDEFINITIONBOOLEANREF = fielValue
				}
			case "SOURCE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SOURCE[identifier].Name = fielValue
				case "SPECOBJECTREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SOURCE[identifier].SPECOBJECTREF = fielValue
				}
			case "SOURCESPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SOURCESPECIFICATION[identifier].Name = fielValue
				case "SPECIFICATIONREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SOURCESPECIFICATION[identifier].SPECIFICATIONREF = fielValue
				}
			case "SPECATTRIBUTES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECATTRIBUTES[identifier].Name = fielValue
				}
			case "SPECHIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECHIERARCHY[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECHIERARCHY[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECHIERARCHY[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECHIERARCHY[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECHIERARCHY[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECIFICATIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONS[identifier].Name = fielValue
				}
			case "SPECIFICATIONTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONTYPE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONTYPE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONTYPE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONTYPE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATIONTYPE[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECIFIEDVALUES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFIEDVALUES[identifier].Name = fielValue
				}
			case "SPECOBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECT[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECT[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECT[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECT[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECT[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECOBJECTS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTS[identifier].Name = fielValue
				}
			case "SPECOBJECTTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTTYPE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTTYPE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTTYPE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTTYPE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECOBJECTTYPE[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECRELATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATION[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATION[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATION[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATION[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATION[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECRELATIONGROUPS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONGROUPS[identifier].Name = fielValue
				}
			case "SPECRELATIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONS[identifier].Name = fielValue
				}
			case "SPECRELATIONTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONTYPE[identifier].Name = fielValue
				case "DESCAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONTYPE[identifier].DESCAttr = fielValue
				case "IDENTIFIERAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONTYPE[identifier].IDENTIFIERAttr = fielValue
				case "LASTCHANGEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONTYPE[identifier].LASTCHANGEAttr = fielValue
				case "LONGNAMEAttr":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECRELATIONTYPE[identifier].LONGNAMEAttr = fielValue
				}
			case "SPECTYPES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECTYPES[identifier].Name = fielValue
				}
			case "TARGET":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TARGET[identifier].Name = fielValue
				case "SPECOBJECTREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TARGET[identifier].SPECOBJECTREF = fielValue
				}
			case "TARGETSPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TARGETSPECIFICATION[identifier].Name = fielValue
				case "SPECIFICATIONREF":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TARGETSPECIFICATION[identifier].SPECIFICATIONREF = fielValue
				}
			case "THEHEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_THEHEADER[identifier].Name = fielValue
				}
			case "TOOLEXTENSIONS":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TOOLEXTENSIONS[identifier].Name = fielValue
				}
			case "VALUES":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_VALUES[identifier].Name = fielValue
				}
			case "XHTMLCONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_XHTMLCONTENT[identifier].Name = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "ALTERNATIVEID":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONBOOLEAN[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONDATE[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].ISEDITABLEAttr = fielValue
				case "MULTIVALUEDAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].MULTIVALUEDAttr = fielValue
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONENUMERATION[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONINTEGER[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONREAL[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONSTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONSTRING[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEDEFINITIONXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].ISEDITABLEAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "DEFAULTVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].DEFAULTVALUE = __gong__map_DEFAULTVALUE[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEDEFINITIONXHTML[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "ATTRIBUTEVALUEBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "THEVALUEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEVALUEBOOLEAN[identifier].THEVALUEAttr = fielValue
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEBOOLEAN[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "ATTRIBUTEVALUEDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEDATE[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "ATTRIBUTEVALUEENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEENUMERATION[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEENUMERATION[identifier].VALUES = __gong__map_VALUES[targetIdentifier]
				}
			case "ATTRIBUTEVALUEINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEINTEGER[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "ATTRIBUTEVALUEREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEREAL[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "ATTRIBUTEVALUESTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUESTRING[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "ATTRIBUTEVALUEXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "ISSIMPLIFIEDAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTEVALUEXHTML[identifier].ISSIMPLIFIEDAttr = fielValue
				case "THEVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEXHTML[identifier].THEVALUE = __gong__map_XHTMLCONTENT[targetIdentifier]
				case "THEORIGINALVALUE":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEXHTML[identifier].THEORIGINALVALUE = __gong__map_XHTMLCONTENT[targetIdentifier]
				case "DEFINITION":
					targetIdentifier := ident.Name
					__gong__map_ATTRIBUTEVALUEXHTML[identifier].DEFINITION = __gong__map_DEFINITION[targetIdentifier]
				}
			case "CHILDREN":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "CORECONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "REQIFCONTENT":
					targetIdentifier := ident.Name
					__gong__map_CORECONTENT[identifier].REQIFCONTENT = __gong__map_REQIFCONTENT[targetIdentifier]
				}
			case "DATATYPEDEFINITIONBOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONBOOLEAN[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPEDEFINITIONDATE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONDATE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPEDEFINITIONENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SPECIFIEDVALUES":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONENUMERATION[identifier].SPECIFIEDVALUES = __gong__map_SPECIFIEDVALUES[targetIdentifier]
				}
			case "DATATYPEDEFINITIONINTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONINTEGER[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPEDEFINITIONREAL":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONREAL[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPEDEFINITIONSTRING":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONSTRING[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPEDEFINITIONXHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_DATATYPEDEFINITIONXHTML[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				}
			case "DATATYPES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DEFAULTVALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "ATTRIBUTEVALUEBOOLEAN":
					targetIdentifier := ident.Name
					__gong__map_DEFAULTVALUE[identifier].ATTRIBUTEVALUEBOOLEAN = __gong__map_ATTRIBUTEVALUEBOOLEAN[targetIdentifier]
				}
			case "DEFINITION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "EDITABLEATTS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "EMBEDDEDVALUE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ENUMVALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_ENUMVALUE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "PROPERTIES":
					targetIdentifier := ident.Name
					__gong__map_ENUMVALUE[identifier].PROPERTIES = __gong__map_PROPERTIES[targetIdentifier]
				}
			case "OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "PROPERTIES":
				switch fieldName {
				// insertion point for field dependant code
				case "EMBEDDEDVALUE":
					targetIdentifier := ident.Name
					__gong__map_PROPERTIES[identifier].EMBEDDEDVALUE = __gong__map_EMBEDDEDVALUE[targetIdentifier]
				}
			case "RELATIONGROUP":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUP[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SOURCESPECIFICATION":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUP[identifier].SOURCESPECIFICATION = __gong__map_SOURCESPECIFICATION[targetIdentifier]
				case "SPECRELATIONS":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUP[identifier].SPECRELATIONS = __gong__map_SPECRELATIONS[targetIdentifier]
				case "TARGETSPECIFICATION":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUP[identifier].TARGETSPECIFICATION = __gong__map_TARGETSPECIFICATION[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUP[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "RELATIONGROUPTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUPTYPE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SPECATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_RELATIONGROUPTYPE[identifier].SPECATTRIBUTES = __gong__map_SPECATTRIBUTES[targetIdentifier]
				}
			case "REQIF":
				switch fieldName {
				// insertion point for field dependant code
				case "HEADER":
					targetIdentifier := ident.Name
					__gong__map_REQIF[identifier].HEADER = __gong__map_THEHEADER[targetIdentifier]
				case "CORECONTENT":
					targetIdentifier := ident.Name
					__gong__map_REQIF[identifier].CORECONTENT = __gong__map_CORECONTENT[targetIdentifier]
				}
			case "REQIFCONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "DATATYPES":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].DATATYPES = __gong__map_DATATYPES[targetIdentifier]
				case "SPECTYPES":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].SPECTYPES = __gong__map_SPECTYPES[targetIdentifier]
				case "SPECOBJECTS":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].SPECOBJECTS = __gong__map_SPECOBJECTS[targetIdentifier]
				case "SPECRELATIONS":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].SPECRELATIONS = __gong__map_SPECRELATIONS[targetIdentifier]
				case "SPECIFICATIONS":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].SPECIFICATIONS = __gong__map_SPECIFICATIONS[targetIdentifier]
				case "SPECRELATIONGROUPS":
					targetIdentifier := ident.Name
					__gong__map_REQIFCONTENT[identifier].SPECRELATIONGROUPS = __gong__map_SPECRELATIONGROUPS[targetIdentifier]
				}
			case "REQIFHEADER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQIFTOOLEXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQTYPE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SOURCE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SOURCESPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECATTRIBUTES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECHIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "ISEDITABLEAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPECHIERARCHY[identifier].ISEDITABLEAttr = fielValue
				case "ISTABLEINTERNALAttr":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPECHIERARCHY[identifier].ISTABLEINTERNALAttr = fielValue
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECHIERARCHY[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "CHILDREN":
					targetIdentifier := ident.Name
					__gong__map_SPECHIERARCHY[identifier].CHILDREN = __gong__map_CHILDREN[targetIdentifier]
				case "EDITABLEATTS":
					targetIdentifier := ident.Name
					__gong__map_SPECHIERARCHY[identifier].EDITABLEATTS = __gong__map_EDITABLEATTS[targetIdentifier]
				case "OBJECT":
					targetIdentifier := ident.Name
					__gong__map_SPECHIERARCHY[identifier].OBJECT = __gong__map_OBJECT[targetIdentifier]
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].VALUES = __gong__map_VALUES[targetIdentifier]
				case "CHILDREN":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].CHILDREN = __gong__map_CHILDREN[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATION[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "SPECIFICATIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECIFICATIONTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATIONTYPE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SPECATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPECIFICATIONTYPE[identifier].SPECATTRIBUTES = __gong__map_SPECATTRIBUTES[targetIdentifier]
				}
			case "SPECIFIEDVALUES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECOBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECOBJECT[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPECOBJECT[identifier].VALUES = __gong__map_VALUES[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPECOBJECT[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "SPECOBJECTS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECOBJECTTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECOBJECTTYPE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SPECATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPECOBJECTTYPE[identifier].SPECATTRIBUTES = __gong__map_SPECATTRIBUTES[targetIdentifier]
				}
			case "SPECRELATION":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATION[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "VALUES":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATION[identifier].VALUES = __gong__map_VALUES[targetIdentifier]
				case "SOURCE":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATION[identifier].SOURCE = __gong__map_SOURCE[targetIdentifier]
				case "TARGET":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATION[identifier].TARGET = __gong__map_TARGET[targetIdentifier]
				case "TYPE":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATION[identifier].TYPE = __gong__map_REQTYPE[targetIdentifier]
				}
			case "SPECRELATIONGROUPS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECRELATIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECRELATIONTYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "ALTERNATIVEID":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATIONTYPE[identifier].ALTERNATIVEID = __gong__map_ALTERNATIVEID[targetIdentifier]
				case "SPECATTRIBUTES":
					targetIdentifier := ident.Name
					__gong__map_SPECRELATIONTYPE[identifier].SPECATTRIBUTES = __gong__map_SPECATTRIBUTES[targetIdentifier]
				}
			case "SPECTYPES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "TARGET":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "TARGETSPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "THEHEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "REQIFHEADER":
					targetIdentifier := ident.Name
					__gong__map_THEHEADER[identifier].REQIFHEADER = __gong__map_REQIFHEADER[targetIdentifier]
				}
			case "TOOLEXTENSIONS":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "VALUES":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "XHTMLCONTENT":
				switch fieldName {
				// insertion point for field dependant code
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "ALTERNATIVEID":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONBOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONDATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONINTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONREAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONSTRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEDEFINITIONXHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEBOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEDATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEINTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEREAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUESTRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTEVALUEXHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CHILDREN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CORECONTENT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONBOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONDATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONINTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONREAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONSTRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPEDEFINITIONXHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DEFAULTVALUE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DEFINITION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "EDITABLEATTS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "EMBEDDEDVALUE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ENUMVALUE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "OBJECT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "PROPERTIES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RELATIONGROUP":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RELATIONGROUPTYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQIF":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQIFCONTENT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQIFHEADER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQIFTOOLEXTENSION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQTYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SOURCE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SOURCESPECIFICATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECATTRIBUTES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECHIERARCHY":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFICATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFICATIONS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFICATIONTYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFIEDVALUES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECOBJECT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECOBJECTS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECOBJECTTYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECRELATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECRELATIONGROUPS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECRELATIONS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECRELATIONTYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECTYPES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TARGET":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TARGETSPECIFICATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "THEHEADER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TOOLEXTENSIONS":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "VALUES":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "XHTMLCONTENT":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
