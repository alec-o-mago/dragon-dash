components {
  id: "basic_green_bullet"
  component: "/components/boss_bullets/basic_green_bullet/basic_green_bullet.script"
  position {
    x: 10.844338
    y: 2.6333764
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"boss_bullet\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"sphere\"\n"
  "  }\n"
  "  data: 12.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"basic_green_bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/battle.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 4.0
    y: 4.0
  }
}
