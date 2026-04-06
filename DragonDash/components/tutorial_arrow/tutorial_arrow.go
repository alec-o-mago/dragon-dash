components {
  id: "tutorial_arrow"
  component: "/components/tutorial_arrow/tutorial_arrow.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tutorial_arrow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/core/core.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 4.0
    y: 4.0
  }
}
