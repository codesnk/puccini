package normal

//
// Artifact
//

type Artifact struct {
	NodeTemplate *NodeTemplate  `json:"-" yaml:"-"`
	Name         string         `json:"-" yaml:"-"`
	Description  string         `json:"description" yaml:"description"`
	Types        Types          `json:"types" yaml:"types"`
	Properties   Constrainables `json:"properties" yaml:"properties"`
	SourcePath   string         `json:"sourcePath" yaml:"sourcePath"`
	TargetPath   string         `json:"targetPath" yaml:"targetPath"`
}

func (self *NodeTemplate) NewArtifact(name string) *Artifact {
	artifact := &Artifact{
		NodeTemplate: self,
		Name:         name,
		Types:        make(Types),
		Properties:   make(Constrainables),
	}
	self.Artifacts[name] = artifact
	return artifact
}

//
// Artifacts
//

type Artifacts map[string]*Artifact
