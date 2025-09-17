components {
  id: "player_bullet"
  component: "/components/player_bullet/player_bullet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player_bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/battle.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 4.0
    y: 4.0
  }
}
