import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { Container, Header, Icon, Image, Menu } from 'semantic-ui-react'
const fixedMenuStyle = {
    backgroundColor: '#fff',
    border: '1px solid #ddd',
    boxShadow: '0px 3px 5px rgba(0, 0, 0, 0.2)',
}
const menuStyle = {
    border: 'none',
    borderRadius: 0,
    boxShadow: 'none',
    marginBottom: '1em',
    marginTop: '4em',
    transition: 'box-shadow 0.5s ease, padding 0.5s ease',
}
const Landing = () => {
    const [menuFixed, setMenuFixed] = useState(false)
    return (
        <div>
            <Menu
                borderless
                fixed={menuFixed ? 'top' : undefined}
                style={menuFixed ? fixedMenuStyle : menuStyle}
            >
                <Container text>
                    <Menu.Item>
                        <Image size='mini' src='/logo.png' />
                    </Menu.Item>
                    <Menu.Item as={Link} to='/' header>Sample</Menu.Item>
                    <Menu.Item as={Link} to='/frontend'>Frontend</Menu.Item>
                    <Menu.Item as={Link} to='/backend'>Backend</Menu.Item>

                    <Menu.Menu position='right'>
                        <Menu.Item as={Link} to='/signin'><Icon name='sign-in' size='big' /></Menu.Item>
                    </Menu.Menu>
                </Container>
            </Menu>
            <Container text>
                <Header as='h2'>Dogs Roles with Humans</Header>
                <p>
                    Domestic dogs inherited complex behaviors, such as bite inhibition, from
                    their wolf ancestors, which would have been pack hunters with complex
                    body language. These sophisticated forms of social cognition and
                    communication may account for their trainability, playfulness, and
                    ability to fit into human households and social situations, and these
                    attributes have given dogs a relationship with humans that has enabled
                    them to become one of the most successful species on the planet today.
      </p>
                <p>
                    The dogs value to early human hunter-gatherers led to them quickly
                    becoming ubiquitous across world cultures. Dogs perform many roles for
                    people, such as hunting, herding, pulling loads, protection, assisting
                    police and military, companionship, and, more recently, aiding
                    handicapped individuals. This impact on human society has given them the
                    nickname mans best friend in the Western world. In some cultures,
                    however, dogs are also a source of meat.
      </p>
            </Container>
            <Container text>
                <Header as='h2'>Dogs Roles with Humans</Header>
                <p>
                    Domestic dogs inherited complex behaviors, such as bite inhibition, from
                    their wolf ancestors, which would have been pack hunters with complex
                    body language. These sophisticated forms of social cognition and
                    communication may account for their trainability, playfulness, and
                    ability to fit into human households and social situations, and these
                    attributes have given dogs a relationship with humans that has enabled
                    them to become one of the most successful species on the planet today.
      </p>
                <p>
                    The dogs value to early human hunter-gatherers led to them quickly
                    becoming ubiquitous across world cultures. Dogs perform many roles for
                    people, such as hunting, herding, pulling loads, protection, assisting
                    police and military, companionship, and, more recently, aiding
                    handicapped individuals. This impact on human society has given them the
                    nickname mans best friend in the Western world. In some cultures,
                    however, dogs are also a source of meat.
      </p>
            </Container>
        </div>
    )
}
export default Landing