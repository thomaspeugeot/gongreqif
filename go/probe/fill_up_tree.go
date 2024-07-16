package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "ALTERNATIVEID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ALTERNATIVEID](probe.stageOfInterest)
			for _alternativeid := range set {
				nodeInstance := (&tree.Node{Name: _alternativeid.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_alternativeid, "ALTERNATIVEID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONBOOLEAN](probe.stageOfInterest)
			for _attributedefinitionboolean := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitionboolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitionboolean, "ATTRIBUTEDEFINITIONBOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONDATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONDATE](probe.stageOfInterest)
			for _attributedefinitiondate := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitiondate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitiondate, "ATTRIBUTEDEFINITIONDATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONENUMERATION](probe.stageOfInterest)
			for _attributedefinitionenumeration := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitionenumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitionenumeration, "ATTRIBUTEDEFINITIONENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONINTEGER](probe.stageOfInterest)
			for _attributedefinitioninteger := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitioninteger.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitioninteger, "ATTRIBUTEDEFINITIONINTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONREAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONREAL](probe.stageOfInterest)
			for _attributedefinitionreal := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitionreal.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitionreal, "ATTRIBUTEDEFINITIONREAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONSTRING](probe.stageOfInterest)
			for _attributedefinitionstring := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitionstring.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitionstring, "ATTRIBUTEDEFINITIONSTRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEDEFINITIONXHTML](probe.stageOfInterest)
			for _attributedefinitionxhtml := range set {
				nodeInstance := (&tree.Node{Name: _attributedefinitionxhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributedefinitionxhtml, "ATTRIBUTEDEFINITIONXHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEBOOLEAN](probe.stageOfInterest)
			for _attributevalueboolean := range set {
				nodeInstance := (&tree.Node{Name: _attributevalueboolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevalueboolean, "ATTRIBUTEVALUEBOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEDATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEDATE](probe.stageOfInterest)
			for _attributevaluedate := range set {
				nodeInstance := (&tree.Node{Name: _attributevaluedate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevaluedate, "ATTRIBUTEVALUEDATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEENUMERATION](probe.stageOfInterest)
			for _attributevalueenumeration := range set {
				nodeInstance := (&tree.Node{Name: _attributevalueenumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevalueenumeration, "ATTRIBUTEVALUEENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEINTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEINTEGER](probe.stageOfInterest)
			for _attributevalueinteger := range set {
				nodeInstance := (&tree.Node{Name: _attributevalueinteger.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevalueinteger, "ATTRIBUTEVALUEINTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEREAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEREAL](probe.stageOfInterest)
			for _attributevaluereal := range set {
				nodeInstance := (&tree.Node{Name: _attributevaluereal.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevaluereal, "ATTRIBUTEVALUEREAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUESTRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUESTRING](probe.stageOfInterest)
			for _attributevaluestring := range set {
				nodeInstance := (&tree.Node{Name: _attributevaluestring.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevaluestring, "ATTRIBUTEVALUESTRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTEVALUEXHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTEVALUEXHTML](probe.stageOfInterest)
			for _attributevaluexhtml := range set {
				nodeInstance := (&tree.Node{Name: _attributevaluexhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributevaluexhtml, "ATTRIBUTEVALUEXHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CHILDREN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CHILDREN](probe.stageOfInterest)
			for _children := range set {
				nodeInstance := (&tree.Node{Name: _children.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_children, "CHILDREN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "CORECONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.CORECONTENT](probe.stageOfInterest)
			for _corecontent := range set {
				nodeInstance := (&tree.Node{Name: _corecontent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_corecontent, "CORECONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONBOOLEAN](probe.stageOfInterest)
			for _datatypedefinitionboolean := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitionboolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitionboolean, "DATATYPEDEFINITIONBOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONDATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONDATE](probe.stageOfInterest)
			for _datatypedefinitiondate := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitiondate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitiondate, "DATATYPEDEFINITIONDATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONENUMERATION](probe.stageOfInterest)
			for _datatypedefinitionenumeration := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitionenumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitionenumeration, "DATATYPEDEFINITIONENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONINTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONINTEGER](probe.stageOfInterest)
			for _datatypedefinitioninteger := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitioninteger.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitioninteger, "DATATYPEDEFINITIONINTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONREAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONREAL](probe.stageOfInterest)
			for _datatypedefinitionreal := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitionreal.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitionreal, "DATATYPEDEFINITIONREAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONSTRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONSTRING](probe.stageOfInterest)
			for _datatypedefinitionstring := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitionstring.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitionstring, "DATATYPEDEFINITIONSTRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPEDEFINITIONXHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPEDEFINITIONXHTML](probe.stageOfInterest)
			for _datatypedefinitionxhtml := range set {
				nodeInstance := (&tree.Node{Name: _datatypedefinitionxhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypedefinitionxhtml, "DATATYPEDEFINITIONXHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPES](probe.stageOfInterest)
			for _datatypes := range set {
				nodeInstance := (&tree.Node{Name: _datatypes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatypes, "DATATYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DEFAULTVALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DEFAULTVALUE](probe.stageOfInterest)
			for _defaultvalue := range set {
				nodeInstance := (&tree.Node{Name: _defaultvalue.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_defaultvalue, "DEFAULTVALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DEFINITION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DEFINITION](probe.stageOfInterest)
			for _definition := range set {
				nodeInstance := (&tree.Node{Name: _definition.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_definition, "DEFINITION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "EDITABLEATTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.EDITABLEATTS](probe.stageOfInterest)
			for _editableatts := range set {
				nodeInstance := (&tree.Node{Name: _editableatts.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_editableatts, "EDITABLEATTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "EMBEDDEDVALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.EMBEDDEDVALUE](probe.stageOfInterest)
			for _embeddedvalue := range set {
				nodeInstance := (&tree.Node{Name: _embeddedvalue.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embeddedvalue, "EMBEDDEDVALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ENUMVALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ENUMVALUE](probe.stageOfInterest)
			for _enumvalue := range set {
				nodeInstance := (&tree.Node{Name: _enumvalue.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enumvalue, "ENUMVALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.OBJECT](probe.stageOfInterest)
			for _object := range set {
				nodeInstance := (&tree.Node{Name: _object.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_object, "OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "PROPERTIES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.PROPERTIES](probe.stageOfInterest)
			for _properties := range set {
				nodeInstance := (&tree.Node{Name: _properties.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_properties, "PROPERTIES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATIONGROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATIONGROUP](probe.stageOfInterest)
			for _relationgroup := range set {
				nodeInstance := (&tree.Node{Name: _relationgroup.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relationgroup, "RELATIONGROUP", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATIONGROUPTYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATIONGROUPTYPE](probe.stageOfInterest)
			for _relationgrouptype := range set {
				nodeInstance := (&tree.Node{Name: _relationgrouptype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relationgrouptype, "RELATIONGROUPTYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQIF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQIF](probe.stageOfInterest)
			for _reqif := range set {
				nodeInstance := (&tree.Node{Name: _reqif.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_reqif, "REQIF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQIFCONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQIFCONTENT](probe.stageOfInterest)
			for _reqifcontent := range set {
				nodeInstance := (&tree.Node{Name: _reqifcontent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_reqifcontent, "REQIFCONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQIFHEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQIFHEADER](probe.stageOfInterest)
			for _reqifheader := range set {
				nodeInstance := (&tree.Node{Name: _reqifheader.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_reqifheader, "REQIFHEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQIFTOOLEXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQIFTOOLEXTENSION](probe.stageOfInterest)
			for _reqiftoolextension := range set {
				nodeInstance := (&tree.Node{Name: _reqiftoolextension.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_reqiftoolextension, "REQIFTOOLEXTENSION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQTYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQTYPE](probe.stageOfInterest)
			for _reqtype := range set {
				nodeInstance := (&tree.Node{Name: _reqtype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_reqtype, "REQTYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SOURCE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SOURCE](probe.stageOfInterest)
			for _source := range set {
				nodeInstance := (&tree.Node{Name: _source.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_source, "SOURCE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SOURCESPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SOURCESPECIFICATION](probe.stageOfInterest)
			for _sourcespecification := range set {
				nodeInstance := (&tree.Node{Name: _sourcespecification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sourcespecification, "SOURCESPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECATTRIBUTES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](probe.stageOfInterest)
			for _specattributes := range set {
				nodeInstance := (&tree.Node{Name: _specattributes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specattributes, "SPECATTRIBUTES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECHIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECHIERARCHY](probe.stageOfInterest)
			for _spechierarchy := range set {
				nodeInstance := (&tree.Node{Name: _spechierarchy.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spechierarchy, "SPECHIERARCHY", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION](probe.stageOfInterest)
			for _specification := range set {
				nodeInstance := (&tree.Node{Name: _specification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification, "SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATIONS](probe.stageOfInterest)
			for _specifications := range set {
				nodeInstance := (&tree.Node{Name: _specifications.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specifications, "SPECIFICATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATIONTYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATIONTYPE](probe.stageOfInterest)
			for _specificationtype := range set {
				nodeInstance := (&tree.Node{Name: _specificationtype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specificationtype, "SPECIFICATIONTYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFIEDVALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFIEDVALUES](probe.stageOfInterest)
			for _specifiedvalues := range set {
				nodeInstance := (&tree.Node{Name: _specifiedvalues.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specifiedvalues, "SPECIFIEDVALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECOBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECOBJECT](probe.stageOfInterest)
			for _specobject := range set {
				nodeInstance := (&tree.Node{Name: _specobject.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specobject, "SPECOBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECOBJECTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECOBJECTS](probe.stageOfInterest)
			for _specobjects := range set {
				nodeInstance := (&tree.Node{Name: _specobjects.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specobjects, "SPECOBJECTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECOBJECTTYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECOBJECTTYPE](probe.stageOfInterest)
			for _specobjecttype := range set {
				nodeInstance := (&tree.Node{Name: _specobjecttype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specobjecttype, "SPECOBJECTTYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECRELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECRELATION](probe.stageOfInterest)
			for _specrelation := range set {
				nodeInstance := (&tree.Node{Name: _specrelation.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specrelation, "SPECRELATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECRELATIONGROUPS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECRELATIONGROUPS](probe.stageOfInterest)
			for _specrelationgroups := range set {
				nodeInstance := (&tree.Node{Name: _specrelationgroups.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specrelationgroups, "SPECRELATIONGROUPS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECRELATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECRELATIONS](probe.stageOfInterest)
			for _specrelations := range set {
				nodeInstance := (&tree.Node{Name: _specrelations.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specrelations, "SPECRELATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECRELATIONTYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECRELATIONTYPE](probe.stageOfInterest)
			for _specrelationtype := range set {
				nodeInstance := (&tree.Node{Name: _specrelationtype.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specrelationtype, "SPECRELATIONTYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECTYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECTYPES](probe.stageOfInterest)
			for _spectypes := range set {
				nodeInstance := (&tree.Node{Name: _spectypes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spectypes, "SPECTYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TARGET":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TARGET](probe.stageOfInterest)
			for _target := range set {
				nodeInstance := (&tree.Node{Name: _target.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_target, "TARGET", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TARGETSPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TARGETSPECIFICATION](probe.stageOfInterest)
			for _targetspecification := range set {
				nodeInstance := (&tree.Node{Name: _targetspecification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_targetspecification, "TARGETSPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "THEHEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.THEHEADER](probe.stageOfInterest)
			for _theheader := range set {
				nodeInstance := (&tree.Node{Name: _theheader.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_theheader, "THEHEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TOOLEXTENSIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TOOLEXTENSIONS](probe.stageOfInterest)
			for _toolextensions := range set {
				nodeInstance := (&tree.Node{Name: _toolextensions.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_toolextensions, "TOOLEXTENSIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.VALUES](probe.stageOfInterest)
			for _values := range set {
				nodeInstance := (&tree.Node{Name: _values.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_values, "VALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "XHTMLCONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.XHTMLCONTENT](probe.stageOfInterest)
			for _xhtmlcontent := range set {
				nodeInstance := (&tree.Node{Name: _xhtmlcontent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtmlcontent, "XHTMLCONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
