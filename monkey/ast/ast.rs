use crate::token::token;
use std::any::Any;


//干...
pub trait Node: std::fmt::Debug {
    fn token_literal(&self) -> String;
    fn node_to_string(&self)->String;
    fn as_any(&self) -> &dyn Any;

    // fn as_super(&self) ->&dyn Node;//还是不要加上为好
}

#[derive(Debug,PartialEq,Clone)]
pub enum ASTNode{
    Program(Program),
    LetStatement(LetStatement),
    BlockStatement(BlockStatement),
    Boolean(Boolean),
    CallExpression(CallExpression),
    ExpressionStatement(ExpressionStatement),
    FuncLiteral(FuncLiteral),
    Identifier(Identifier),
    IfExpression(IfExpression),
    InfixExpression(InfixExpression),
    IntegerLiteral(IntegerLiteral),
    PrefixExpression(PrefixExpression),
    ReturnStatement(ReturnStatement),
    None,
}


impl ASTNode{
    pub fn new()-> Self{
        ASTNode::None
    }
    pub fn is_none(&self) -> bool{
        use self::ASTNode::*;
        match self{
            None=>true,
            _=>false,
        }
    }
    pub fn is_some(&self) ->bool{
        !self.is_none()
    }
}

// impl<T> Node for T where T: Node{
//     fn as_super(&self) ->&dyn Node {
//         self
//     }
// } 

#[derive(Debug,PartialEq,Clone)]
pub struct Program {
    pub statements: Vec<Box<ASTNode>>,
}


impl Program {
    pub fn new()->Self{
        Program{ statements:vec![] }
    }
    
    // fn node_to_string(&self)->String{
    //     let mut s = String::new();
    //     for i in &self.statements {
    //         match i{
    //             ASTNode::None => s.push_str(&stmt.node_to_string()),
    //             _=>  s.push_str(""),
    //         }
    //     }
    //     s
    // }
}

// impl  Node for Program {
//     fn token_literal(&self) ->String {
//         if self.statements.len()>0 {
//             return match self.statements[0]{
//                 Some(ref stmt) => stmt.node_to_string(),
//                 None =>"".to_string(),
//             }
            
//         }
//         "".to_string()
//     }
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
//         for i in &self.statements {
//             match i{
//                 Some(stmt)=> s.push_str(&stmt.node_to_string()),
//                 _=>  s.push_str(""),
//             }
//         }
//         s
//     }
//     fn as_any(&self)->&dyn Any{
//         self
//     }
// }

#[derive(Debug,PartialEq,Clone)]
pub struct LetStatement{
    pub token: token::Token,
    pub name: Box<Identifier>,
    pub value: Box<ASTNode>,
}

impl LetStatement{
    pub fn new()->LetStatement{
        LetStatement{
            token: token::Token::default(),
            name:Box::new(Identifier::default()),
            value:Box::new(ASTNode::None),
        }
    }
    pub fn is_none(&self)->bool{
        self.value.is_none()
    }
}

// impl Node for LetStatement{
//     fn statement_node(&self) {}
    
// }

// impl Node for LetStatement{
//     fn token_literal(&self)->String{ self.token.Literal.clone()}
//     fn node_to_string(&self)->String{ 
//         let mut s = String::new();
//         s.push_str(&self.token_literal());
//         s.push_str(" ");
//         s.push_str(&self.name.node_to_string());
//         s.push_str(" = ");

//         if let Some(ref value) = self.value{
//             s.push_str(&value.node_to_string());
//         }
//         s.push_str(";");
//         s
//     }
    
//     fn as_any(&self)->&dyn Any{
//         self
//     }

// }


#[derive(Debug,PartialEq,Clone)]
pub struct Identifier{
    pub token: token::Token,
    pub value: String,
}

impl Identifier{
    pub fn new(token: token::Token, value: String) -> Self{
        Self{
            token,value,
        }
    }
    
    pub fn default()->Self{
        Identifier{
            token: token::Token::default(), 
            value: String::default()
        }
    }
}


// impl Node for Identifier{
//     fn token_literal(&self) ->String {self.token.Literal.clone()}
//     fn node_to_string(&self)->String{
//         self.value.to_string()
//     }
//     fn as_any(&self)->&dyn Any{
//         self
//     }
// }


#[derive(Debug,PartialEq,Clone)]
pub struct ReturnStatement{
    pub token: token::Token,
    pub return_value: Box<ASTNode>,
}

impl ReturnStatement{
    pub fn new(token: token::Token)->Self{
        ReturnStatement{token: token, return_value:Box::new(ASTNode::None),}
    }
}

// impl Node for ReturnStatement{
//     fn token_literal(&self) ->String {self.token.Literal.clone()}
//     fn as_any(&self)->&dyn Any{
//         self
//     }

//     fn node_to_string(&self)->String{ 
//         let mut s = String::new();
//         s.push_str(&self.token_literal());
//         s.push_str(" ");
//         if let Some(ref value) = self.return_value{
//             s.push_str(&value.node_to_string());
//         }
//         s.push_str(";");
//         s
//     }
// }


//Expression 

#[derive(Debug,PartialEq,Clone)]
pub struct ExpressionStatement{
    token: token::Token,
    pub expression: Box<ASTNode>,
}

impl ExpressionStatement{
    pub fn new(token: token::Token)->Self{
        ExpressionStatement{token: token, 
            expression:Box::new(ASTNode::None),
        }
    }
}

// impl Node for ExpressionStatement{
//     fn token_literal(&self) ->String {self.token.Literal.clone() }
//     fn node_to_string(&self)->String{
//         match self.expression{
//             Some(ref expression) => expression.node_to_string(),
//             _=>"".to_string(),
//         }
//     }
//     fn as_any(&self)->&dyn Any{self}
// }


#[derive(Debug,PartialEq,Clone)]
pub struct IntegerLiteral{
    token: token::Token,
    pub value: i64,
}
impl IntegerLiteral{
    pub fn new(token: token::Token)->Self{
        IntegerLiteral{token: token, value:0,}
    }
}

// impl Node for IntegerLiteral{
//     fn token_literal(&self) ->String {self.token.Literal.clone()}
//     fn node_to_string(&self)->String{self.token.Literal.clone()}
//     fn as_any(&self)->&dyn Any{self}
// }



#[derive(Debug,PartialEq,Clone)]
pub struct PrefixExpression{
    token: token::Token,
    pub operator: String,
    pub right: Box<ASTNode>,
}

impl PrefixExpression{
    pub fn new<S: Into<String>>(token: token::Token,operator: S)->Self{
        PrefixExpression{
            token: token,
            operator: operator.into(),
            right: Box::new(ASTNode::None),
        }
    }
}

// impl Node for PrefixExpression{
//     fn as_any(&self)->&dyn Any{self}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
//         s.push('(');
        
//         s.push_str(&self.operator.clone());
//         let tmp = match self.right{
//             Some(ref right) => right.node_to_string().clone(),
//             _=>"".to_string(),
//         };
//         s.push_str(&tmp);
//         s.push(')');
//         s
//     }
//     fn token_literal(&self)->String{self.token.Literal.clone()}
// }


#[derive(Debug,PartialEq,Clone)]
pub struct InfixExpression{
    token: token::Token,
    pub left: Box<ASTNode>,
    pub right: Box<ASTNode>,
    pub operator: String,
}

impl InfixExpression{
    pub fn new<S: Into<String>>(token: token::Token, operator:S)->Self{
        InfixExpression{
            token: token,
            operator: operator.into(),
            left: Box::new(ASTNode::None),
            right:Box::new(ASTNode::None),
        }
    }
}

// impl Node for InfixExpression{
//     fn as_any(&self)->&dyn Any{self}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
//         s.push('(');
//         let tmpl = match self.left{
//             Some(ref left) => left.node_to_string().clone(),
//             _=>"".to_string(),
//         };
//         s.push_str(&tmpl);

//         s.push(' ');
//         s.push_str(&self.operator.clone());
//         s.push(' ');
        
//         let tmp = match self.right{
//             Some(ref right) => right.node_to_string().clone(),
//             _=>"".to_string(),
//         };
//         s.push_str(&tmp);
//         s.push(')');
//         s
//     }
//     fn token_literal(&self)->String{self.token.Literal.clone()}
// }


#[derive(Debug,PartialEq,Clone)]
pub struct Boolean{
    pub token: token::Token, 
    pub value: bool,
}
impl Boolean{
    pub fn new(token: token::Token, value: bool) ->Self{
        Boolean{
                token: token,
                value:value,
        }
    }
}

// impl Node for Boolean{
//     fn as_any(&self)->&dyn Any{self}
//     fn token_literal(&self)->String{self.token.Literal.clone()}
//     fn node_to_string(&self)->String{self.token.Literal.clone()}
// }



#[derive(Debug,PartialEq,Clone)]
pub struct IfExpression{
    token: token::Token,
    pub condition: Box<ASTNode>,
    pub consequence: Box<ASTNode>,
    pub alternative: Box<ASTNode>,

}


impl IfExpression{
    pub fn new(token: token::Token)->Self{
        IfExpression{
            token: token,
            condition: Box::new(ASTNode::None),
            consequence: Box::new(ASTNode::None),
            alternative: Box::new(ASTNode::None),
        }
    }
}

// impl Node for IfExpression{
//     fn as_any(&self)->&dyn Any{self}
//     fn token_literal(&self)->String{self.token.Literal.clone()}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();

//         s.push_str("if");
//         if let Some(ref value) = self.condition{
//             s.push_str(&value.node_to_string());
//         }
//         s.push_str(" ");
//         if let Some(ref value) = self.consequence{
//             s.push_str(&value.node_to_string());
//         }
//         if let Some(ref value) = self.alternative{
//             s.push_str("else ");
//             s.push_str(&value.node_to_string());
//         }
//         s
//     }
// }


#[derive(Debug,PartialEq,Clone)]
pub struct BlockStatement{
    token: token::Token,
    pub statements: Vec<Box<ASTNode>>,
}

impl BlockStatement{
    pub fn new(token: token::Token)->Self{
        BlockStatement{
                token: token, 
                statements:vec![],
            }
    }

    // pub to_string(&self)->String{

    // }
}

// impl Node for BlockStatement{
//     fn as_any(&self)->&dyn Any{self}
//     fn token_literal(&self)->String{self.token.Literal.clone()}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
//         for i in &self.statements{
//             if let Some(ref value) = i{
//                 s.push_str(&value.node_to_string());
//             }
//         }
//         s
//     }
// }


#[derive(Debug,PartialEq,Clone)]
pub struct FuncLiteral{
    token: token::Token,
    pub params: Vec<Option<Box<Identifier>>>,
    pub body: Box<ASTNode>,
}

impl FuncLiteral{
    pub fn new(token: token::Token)->Self{
        FuncLiteral{token: token, params: Vec::new(), body: Box::new(ASTNode::None),}
    }

}

// impl Node for FuncLiteral{
//     fn as_any(&self)->&dyn Any{self}
//     fn token_literal(&self)->String{self.token.Literal.clone()}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
        

//         s.push_str(&self.token_literal());
//         s.push('(');
//         for i in &self.params{
//             if let Some(ref value) = i{
//                 s.push_str(&value.node_to_string());
//                 s.push(',');
//             }
//         }
//         s.push(')');
//         if let Some(ref value) = self.body{
//             s.push_str(&value.node_to_string());
//         }
        
//         s
//     }
// }


#[derive(Debug,PartialEq,Clone)]
pub struct CallExpression{
    token: token::Token,
    pub func: Box<ASTNode>,
    pub args: Vec<Box<ASTNode>>,
}

impl CallExpression{
    pub fn new(token: token::Token)->Self{
        CallExpression{
            token:token,
            func: Box::new(ASTNode::None),
            args: Vec::new(),
        }
    }
}


// impl Node for CallExpression{
//     fn as_any(&self)->&dyn Any{self}
//     fn token_literal(&self)->String{self.token.Literal.clone()}
//     fn node_to_string(&self)->String{
//         let mut s = String::new();
        
//         if let Some(ref value) = self.func{
//             s.push_str(&value.node_to_string());
//         }
//         s.push('(');
//         for i in &self.args{
//             if let Some(ref value) = i{
//                 s.push_str(&value.node_to_string());
//                 if self.args.len() > 1 {
//                     s.push(',');
//                 }
                
//             }
//         }
//         s.push(')');
        
//         s
//     }
// }
