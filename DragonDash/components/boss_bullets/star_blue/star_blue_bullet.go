components {
  id: "star_blue_bullet"
  component: "/components/boss_bullets/star_blue/star_blue_bullet.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"boss_bullets\"\n"
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
  "  data: 21.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"star_boss_blue_bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/core/core.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.8
  }
  scale {
    x: 4.0
    y: 4.0
  }
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/components/boss_bullets/blue/blue_bullet.go\"\n"
  ""
}
