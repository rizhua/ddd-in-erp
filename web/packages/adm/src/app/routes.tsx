import { RouteObject, Outlet } from 'react-router-dom';

import {common, domain, layout} from '@/component';

const routes: RouteObject[] = [
    {
        path: '/',
        element: <layout.Console />,
        children: [{
            path: 'bundle',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.bundle.List />
                },
            ]
        }, {
            path: 'customer',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.customer.List />
                },
            ]
        }, {
            path: 'structure',
            element: <Outlet />,
            children: [
                {
                    path: 'dept',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.structure.dept.List />
                        },
                    ],
                }, {
                    path: 'org',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.structure.org.List />
                        },
                    ]
                },
            ]
        }, {
            path: 'emp',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.emp.List />
                }, {
                    path: 'invite',
                    element: <domain.emp.Invite />
                }
            ]
        }, {
            path: 'market',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.market.List />
                }
            ]
        }, {
            path: 'node',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.node.List />
                },
            ]
        }, {
            path: 'notice',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.notice.List />
                },
                {
                    path: 'edit/:id',
                    element: <domain.notice.Edit />
                },
            ]
        }, {
            path: 'product',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.product.List />,
                }, {
                    path: 'publish',
                    element: <domain.product.Publish />,
                }, {
                    path: 'attribute',
                    element: <domain.product.attribute.List />,
                }, {
                    path: 'comment',
                    element: <domain.product.Comment />,
                }, {
                    path: 'brand',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.product.brand.List />
                        }
                    ]
                }, {
                    path: 'category',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.product.category.List />
                        },
                    ]
                }, {
                    path: 'agent',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.product.agent.List />
                        }
                    ]
                }, {
                    path: 'position',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.product.position.List />
                        }
                    ]
                }
            ]
        }, {
            path: 'purchase',
            element: <Outlet />,
            children: [
                {
                    path: 'plan',
                    element: <domain.purchase.plan.List />,
                }
            ]
        }, {
            path: 'role',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.role.List />
                },
            ]
        }, {
            path: 'space',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.space.List />,
                }, {
                    path: 'goods',
                    element: <Outlet />,
                    children: [
                        {
                            path: '',
                            element: <domain.space.goods.List />
                        },
                    ]
                }
            ]
        }, {
            path: 'user',
            element: <Outlet />,
            children: [
                {
                    path: '',
                    element: <domain.user.List />
                }, {
                    path: 'attest',
                    element: <domain.user.Attest />
                }, {
                    path: 'profile',
                    element: <domain.user.Profile />
                }, {
                    path: 'security',
                    element: <domain.user.Security />
                }, {
                    path: 'subscribe',
                    element: <domain.user.Subcribe />
                },
            ]
        }]
    },{
        path: '/auth',
        element: <Outlet />,
        children: [
            {
                path: 'signIn',
                element: <domain.auth.SignIn />
            }, {
                path: 'signUp',
                element: <domain.auth.SignUp />
            },
        ]
    }, {
        path: '/password',
        element: <Outlet />,
        children: [
            {
                path: 'forget',
                element: <domain.password.Forget />
            }, {
                path: 'reset',
                element: <domain.password.Reset />
            },
        ]
    }, {
        path: '*',
        element: <common.NotFound />
    }
];

export default routes;